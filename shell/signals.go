package shell

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

const (
	// PromptSIGINT defines the string to write when ctrl+c is pressed
	PromptSIGINT = "^C"

	// PromptSIGQUIT defines the string to write when ctrl+\ is pressed
	PromptSIGQUIT = "^\\"

	// PromptEOF defines the string to write when ctrl+d is pressed
	PromptEOF = "^D"
)

func sigtstp() {
	show, err := lang.ShellProcess.Config.Get("shell", "show-stop-status", types.Boolean)
	if err != nil {
		show = false
	}
	if !show.(bool) {
		return
	}

	/*defer func() {
		if debug.Enabled {
			return
		}
		if r := recover(); r != nil {
			return
		}
	}()*/

	p := lang.ForegroundProc
	stdinR, stdinW := p.Stdin.Stats()
	stdoutR, stdoutW := p.Stdout.Stats()
	stderrR, stderrW := p.Stderr.Stats()
	pipeStatus := fmt.Sprintf(
		"\nSTDIN:  %s read /%s written\nSTDOUT: %s read / %s written\nSTDERR: %s read /%s written",
		utils.HumanBytes(stdinR), utils.HumanBytes(stdinW),
		utils.HumanBytes(stdoutR), utils.HumanBytes(stdoutW),
		utils.HumanBytes(stderrR), utils.HumanBytes(stderrW),
	)
	lang.ShellProcess.Stderr.Writeln([]byte(pipeStatus))

	if p.Exec.Pid != 0 {
		block, err := lang.ShellProcess.Config.Get("shell", "stop-status-func", types.CodeBlock)
		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			return
		}

		//branch := lang.ShellProcess.BranchFID()
		//defer branch.Close()
		//branch.Variables.Set("PID", lang.ForegroundProc.Exec.Pid, types.Integer)
		//_, err = lang.RunBlockExistingConfigSpace([]rune(block.(string)), nil, lang.ShellProcess.Stdout, lang.ShellProcess.Stderr, branch.Process)

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN)
		fork.Variables.Set("PID", lang.ForegroundProc.Exec.Pid, types.Integer)
		fork.Execute([]rune(block.(string)))

		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		}

		lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
			"FID %d has been stopped. Use `fg %d` / `bg %d` to manage the FID or `jobs` or `fid-list` to see a list of processes running on this shell.",
			p.Id, p.Id, p.Id,
		)))

		p.State = state.Stopped

		go ShowPrompt()

	} else {
		lang.ShellProcess.Stderr.Write([]byte("(murex functions don't currently support being stopped)"))
	}

}

func sigint(interactive bool) {
	//os.Stderr.WriteString(PromptSIGINT)
	sigterm(interactive)
}

func sigterm(interactive bool) {
	if interactive {
		if lang.ForegroundProc != nil && lang.ForegroundProc.Kill != nil {
			lang.ForegroundProc.Kill()
		}

	} else {
		os.Exit(0)
	}
}

func sigquit(interactive bool) {
	if interactive {
		//os.Stderr.WriteString(PromptSIGQUIT)
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)

		fids := lang.GlobalFIDs.ListAll()
		for _, p := range fids {
			if p.Kill != nil /*&& !p.HasTerminated()*/ {
				procName := p.Name
				procParam, _ := p.Parameters.String(0)
				if p.Name == "exec" {
					procName = procParam
					procParam, _ = p.Parameters.String(1)
				}
				if len(procParam) > 10 {
					procParam = procParam[:10]
				}
				lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf("!!! Sending kill signal to fid %d: %s %s !!!", p.Id, procName, procParam)))
				p.Kill()
			}
		}

		lang.ShellProcess.Stderr.Writeln([]byte("!!! Starting new prompt !!!"))
		go ShowPrompt()

	} else {
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)
		os.Exit(2)
	}
}
