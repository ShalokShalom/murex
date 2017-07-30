package lang

import (
	"errors"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

var ShellExitNum int // for when running murex in interactive shell mode

func ProcessNewBlock(block []rune, stdin, stdout, stderr streams.Io, caller *proc.Process) (exitNum int, err error) {
	container := new(proc.Process)
	container.State = state.MemAllocated
	container.IsBackground = caller.IsBackground
	container.Name = caller.Name
	//container.Parent = nil
	container.Parent = caller
	container.Id = caller.Id
	if caller.Name == "shell" {
		container.ExitNum = ShellExitNum
	}

	if stdin != nil {
		container.Stdin = stdin
	} else {
		container.Stdin = streams.NewStdin()
		container.Stdin.SetDataType(types.Null)
		container.Stdin.Close()
	}

	if stdout != nil {
		container.Stdout = stdout
	} else {
		container.Stdout = new(streams.TermOut)
	}
	container.Stdout.MakeParent()

	if stderr != nil {
		container.Stderr = stderr
	} else {
		container.Stderr = new(streams.TermErr)
	}
	container.Stderr.MakeParent()

	tree, pErr := ParseBlock(block)
	if pErr.Code != 0 {
		container.Stderr.Writeln([]byte(pErr.Message))
		debug.Json("ParseBlock returned:", pErr)
		err = errors.New(pErr.Message)
		return 1, err
	}

	compile(&tree, container)

	// Support for different run modes:
	switch {
	case container.Name == "try":
		exitNum = runHyperSensitive(&tree)
	default:
		exitNum = runNormal(&tree)
		//exitNum = runHyperSensitive(&tree)
	}

	// This will just unlock the parent lock. Stdxxx.Close() will still have to be called.
	container.Stdout.UnmakeParent()
	container.Stderr.UnmakeParent()

	debug.Json("Finished running &tree", tree)
	return
}
