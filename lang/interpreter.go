package lang

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
)

func compile(tree *astNodes, parent *proc.Process) (procs []proc.Process) {
	procs = make([]proc.Process, len(*tree))

	for i := range *tree {
		procs[i].State = state.MemAllocated
		procs[i].Name = (*tree)[i].Name
		procs[i].IsMethod = (*tree)[i].Method
		procs[i].IsBackground = parent.IsBackground
		procs[i].Parent = parent
		procs[i].Scope = parent.Scope
		procs[i].WaitForTermination = make(chan bool)
		procs[i].RunMode = parent.RunMode
		procs[i].Config = parent.Config
		procs[i].Variables = proc.ReferenceVariables(parent.Variables)

		procs[i].Parameters.SetTokens((*tree)[i].ParamTokens)

		procs[i].FidTree = make([]int, len(parent.Parent.FidTree))
		copy(procs[i].FidTree, parent.Parent.FidTree)

		if (*tree)[i].LineNumber == 0 {
			procs[i].ColNumber = (*tree)[i].ColNumber + parent.ColNumber
		} else {
			procs[i].ColNumber = (*tree)[i].ColNumber
		}

		if parent.Id == 0 {
			procs[i].LineNumber = (*tree)[i].LineNumber + parent.LineNumber + 1
		} else {
			procs[i].LineNumber = (*tree)[i].LineNumber + parent.LineNumber
		}

		// Define previous and next processes:
		switch {
		case i == 0:
			// first
			procs[0].Previous = parent
			if i == len(*tree)-1 {
				procs[0].Next = parent
			} else {
				procs[0].Next = &procs[1]
			}

		case i == len(*tree)-1:
			// last
			procs[i].Previous = &procs[i-1]
			procs[i].Next = parent

		case i > 0:
			// everything in the middle
			procs[i].Previous = &procs[i-1]
			procs[i].Next = &procs[i+1]

		default:
			// This condition should never happen,
			// but lets but a default catch and stack trace in just in case.
			panic("Failed in an unexpected way: Compile()->switch{default}")
		}

		// Define stdin interface:
		switch {
		case i == 0:
			// first
			procs[0].Stdin = parent.Stdin

		case (*tree)[i].NewChain:
			// new chain
			procs[i].Stdin = streams.NewStdin()
		}

		// Define stdout / stderr interfaces:
		switch {
		case (*tree)[i].PipeOut:
			procs[i+1].Stdin = streams.NewStdin()
			procs[i].Stdout = procs[i].Next.Stdin
			procs[i].Stderr = procs[i].Parent.Stderr

		case (*tree)[i].PipeErr:
			procs[i+1].Stdin = streams.NewStdin()
			procs[i].Stdout = procs[i].Parent.Stdout
			procs[i].Stderr = procs[i].Next.Stdin

		default:
			procs[i].Stdout = procs[i].Parent.Stdout
			procs[i].Stderr = procs[i].Parent.Stderr
		}

	}

	for i := range *tree {
		createProcess(&procs[i], !(*tree)[i].NewChain)
	}

	return
}

// `evil` - Only use this if you are not concerned about STDERR nor exit number.
func runModeEvil(procs []proc.Process) int {
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := range procs {

		if i > 0 {
			if !procs[i].IsMethod {
				waitProcess(&procs[i-1])
			} else {
				go waitProcess(&procs[i-1])
			}
		}

		/*if procs[i].Name == "break" {
			exitNum, _ := procs[i].Parameters.Int(0)
			return exitNum
		}*/
		procs[i].Stderr = new(streams.Null)
		go executeProcess(&procs[i])
	}

	waitProcess(&procs[len(procs)-1])
	return 0
}

func runModeNormal(procs []proc.Process) (exitNum int) {
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := range procs {

		if i > 0 {
			if !procs[i].IsMethod {
				waitProcess(&procs[i-1])
			} else {
				go waitProcess(&procs[i-1])
			}
		}

		/*if procs[i].Name == "break" {
			exitNum, _ = procs[i].Parameters.Int(0)
			return
		}*/
		go executeProcess(&procs[i])
	}

	waitProcess(&procs[len(procs)-1])
	exitNum = procs[len(procs)-1].ExitNum
	return
}

// `try` - Last process in each pipe is checked.
func runModeTry(procs []proc.Process) (exitNum int) {
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := range procs {
		if i > 0 {
			if !procs[i].IsMethod {
				waitProcess(&procs[i-1])
				exitNum = procs[i-1].ExitNum
				outSize, _ := procs[i-1].Stdout.Stats()
				errSize, _ := procs[i-1].Stderr.Stats()

				if exitNum == 0 && errSize > outSize {
					exitNum = 1
				}

				if exitNum > 0 {
					for ; i < len(procs); i++ {
						procs[i].Stdout.Close()
						procs[i].Stderr.Close()
						proc.GlobalFIDs.Deregister(procs[i].Id)
						procs[i].State = state.AwaitingGC
					}
					return
				}

			} else {
				go waitProcess(&procs[i-1])
			}
		}

		/*if procs[i].Name == "break" {
			exitNum, _ = procs[i].Parameters.Int(0)
			return
		}*/
		go executeProcess(&procs[i])
	}

	last := len(procs) - 1
	waitProcess(&procs[last])
	exitNum = procs[last].ExitNum
	outSize, _ := procs[last].Stdout.Stats()
	errSize, _ := procs[last].Stderr.Stats()

	if exitNum == 0 && errSize > outSize {
		exitNum = 1
	}

	return
}

// `trypipe` - Each process in the pipeline is tried sequentially. Breaks parallelisation.
func runModeTryPipe(procs []proc.Process) (exitNum int) {
	//debug.Log("Entering run mode `trypipe`")
	if len(procs) == 0 {
		return 1
	}

	procs[0].Previous.SetTerminatedState(true)

	for i := range procs {
		/*if procs[i].Name == "break" {
			exitNum, _ = procs[i].Parameters.Int(0)
			return
		}*/
		go executeProcess(&procs[i])
		waitProcess(&procs[i])

		exitNum = procs[i].ExitNum
		outSize, _ := procs[i].Stdout.Stats()
		errSize, _ := procs[i].Stderr.Stats()

		if exitNum == 0 && errSize > outSize {
			exitNum = 1
		}

		if exitNum > 0 {
			for i++; i < len(procs); i++ {
				procs[i].Stdout.Close()
				procs[i].Stderr.Close()
				proc.GlobalFIDs.Deregister(procs[i].Id)
				procs[i].State = state.AwaitingGC
			}
			return
		}
	}

	return
}
