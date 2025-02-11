package shell

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/ansititle"
	"github.com/lmorg/murex/utils/cd"
	"github.com/lmorg/murex/utils/cd/cache"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/counter"
	"github.com/lmorg/murex/utils/readline"
	"github.com/lmorg/murex/utils/spellcheck"
)

var (
	// Prompt is the readline instance
	Prompt = readline.NewInstance()

	// PromptId is an custom defined ID for each prompt Goprocess so we don't
	// accidentally end up with multiple prompts running
	PromptId = new(counter.MutexCounter)

	// Events is a callback for onPrompt events
	Events func(string, []rune)
)

func callEvents(interrupt string, cmdLine []rune) {
	if Events == nil {
		return
	}
	Events(interrupt, cmdLine)
}

// Start the interactive shell
func Start() {
	if debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				tty.Stderr.WriteString(fmt.Sprintln("Panic caught:", r))
				Start()
			}
		}()
	}

	lang.ShellProcess.StartTime = time.Now()

	// disable this for Darwin (macOS) because the messages it pops up might
	// spook many macOS users.
	if runtime.GOOS != "darwin" {
		go cache.GatherFileCompletions(".")
	}

	v, err := lang.ShellProcess.Config.Get("shell", "pre-cache-hint-summaries", types.String)
	if err != nil {
		v = ""
	}
	if v.(string) == types.TrueString || v.(string) == "on-start" {
		go autocomplete.CacheHints()
	}

	lang.Interactive = true
	Prompt.TempDirectory = consts.TempDir

	definePromptHistory()
	Prompt.AutocompleteHistory = autocompleteHistoryLine

	SignalHandler(true)

	v, err = lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}
	Prompt.MaxTabCompleterRows = v.(int)

	pwd, _ := lang.ShellProcess.Config.Get("shell", "start-directory", types.String)
	pwd = strings.TrimSpace(pwd.(string))
	if pwd != "" {
		err := cd.Chdir(lang.ShellProcess, pwd.(string))
		if err != nil {
			tty.Stderr.WriteString(err.Error())
		}
	}

	ShowPrompt()

	noQuit := make(chan int)
	<-noQuit
}

// ShowPrompt display's the shell command line prompt
func ShowPrompt() {
	if !lang.Interactive {
		panic("shell.ShowPrompt() called before initialising prompt with shell.Start()")
	}

	var (
		thisProc = PromptId.Add()
		nLines   = 1
		merged   string
		block    []rune
	)

	Prompt.GetMultiLine = func(r []rune) []rune {
		var multiLine []rune
		if len(block) == 0 {
			multiLine = r
		} else {
			multiLine = append(append(block, []rune(utils.NewLineString)...), r...)
		}

		expanded, err := history.ExpandVariables(multiLine, Prompt)
		if err != nil {
			expanded = multiLine
		}
		return expanded
	}

	for {
		v, err := lang.ShellProcess.Config.Get("proc", "echo-tmux", types.Boolean)
		if tmux, ok := v.(bool); ok && err == nil && tmux {
			ansititle.Tmux([]byte(app.Name))
		}

		signalRegister(true)

		setPromptHistory()
		Prompt.TabCompleter = tabCompletion
		Prompt.SyntaxCompleter = syntaxCompletion
		Prompt.DelayedSyntaxWorker = Spellchecker
		Prompt.HistoryAutoWrite = false

		if tty.Enabled() {
			Prompt.ScreenRefresh = tty.BufferGet
		}

		getSyntaxHighlighting()
		getHintTextEnabled()
		getHintTextFormatting()
		getPreviewSettings()
		cachedHintText = []rune{}
		var prompt []byte

		if nLines > 1 {
			prompt = getMultilinePrompt(nLines)
		} else {
			callEvents("before", nil)
			block = []rune{}
			prompt = getPrompt()
			writeTitlebar()
		}

		if tty.MissingCrLf() {
			tty.WriteCrLf()
		}

		Prompt.SetPrompt(string(prompt))

		line, err := Prompt.Readline()
		if err != nil {
			switch err {
			case readline.CtrlC:
				merged = ""
				nLines = 1
				fmt.Fprintln(tty.Stdout, PromptSIGINT)
				callEvents("cancel", nil)
				continue

			case readline.EOF:
				fmt.Fprintln(tty.Stdout, utils.NewLineString)
				callEvents("eof", nil)
				lang.Exit(0)

			default:
				fmt.Fprint(os.Stderr, err.Error())
			}
		}

		switch {
		case nLines > 1:
			block = append(block, []rune(utils.NewLineString+line)...)
		default:
			block = []rune(line)
		}

		expanded, err := history.ExpandVariables(block, Prompt)
		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			merged = ""
			nLines = 1
			continue
		}

		if string(expanded) != string(block) {
			tty.Stdout.WriteString(ansi.ExpandConsts("{GREEN}") + string(expanded) + ansi.ExpandConsts("{RESET}") + utils.NewLineString)
		}

		pt, _ := parse(block)
		switch {
		case pt.NestedBlock > 0:
			nLines++
			merged += line + `^\n`

		case pt.Escaped:
			nLines++
			merged += strings.TrimSpace(line[:len(line)-1]) + " "

		case pt.QuoteSingle, pt.QuoteBrace > 0:
			nLines++
			merged += line + `^\n`

		case pt.QuoteDouble:
			nLines++
			merged += line + `\n`

		case len(block) == 0:
			continue

		default:
			merged += line
			mergedExp, err := history.ExpandVariablesInLine([]rune(merged), Prompt)
			if err == nil {
				merged = string(mergedExp)
			}

			macroFind, macroReplace, err := getMacroVars(merged)
			if err != nil {
				nLines = 1
				merged = ""
				continue
			}

			if len(macroFind) > 0 {
				expanded = []rune(expandMacroVars(string(expanded), macroFind, macroReplace))
			}

			_, err = Prompt.History.Write(merged)
			if err != nil {
				fmt.Fprintf(tty.Stdout, ansi.ExpandConsts("{RED}Error: cannot write history file: %s{RESET}\n"), err.Error())
			}

			nLines = 1
			merged = ""

			callEvents("after", block)

			fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN)
			fork.FileRef = ref.NewModule(app.ShellModule)
			fork.Stderr = term.NewErr(ansi.IsAllowed())
			fork.PromptId = thisProc
			fork.CCEvent = lang.ShellProcess.CCEvent
			fork.CCExists = lang.ShellProcess.CCExists
			lang.ShellExitNum, err = fork.Execute(expanded)
			if err != nil {
				fmt.Fprintln(tty.Stdout, ansi.ExpandConsts(fmt.Sprintf("{RED}%v{RESET}", err)))
			}

			if tty.MissingCrLf() {
				tty.WriteCrLf()
			}

			if PromptId.NotEqual(thisProc) {
				return
			}
		}
	}
}

var rxMacroVar = regexp.MustCompile(`(\^\$[-_a-zA-Z0-9]+)`)

func getMacroVars(s string) ([]string, []string, error) {
	var err error

	if !rxMacroVar.MatchString(s) {
		return nil, nil, nil
	}

	assigned := make(map[string]bool)

	match := rxMacroVar.FindAllString(s, -1)
	vars := make([]string, len(match))
	for i := range match {
		if assigned[match[i]] {
			continue
		}

		for {
			rl := readline.NewInstance()
			rl.SetPrompt(ansi.ExpandConsts(fmt.Sprintf(
				"{YELLOW}Enter value for: {RED}%s{YELLOW}? {RESET}", match[i][2:],
			)))
			rl.History = new(readline.NullHistory)
			vars[i], err = rl.Readline()
			if err != nil {
				return nil, nil, err
			}
			if vars[i] != "" {
				break
			}
			tty.Stderr.WriteString(ansi.ExpandConsts("{RED}Cannot use zero length strings. Please enter a value or press CTRL+C to cancel.{RESET}\n"))
		}
		assigned[match[i]] = true
	}

	return match, vars, nil
}

func expandMacroVars(s string, match, vars []string) string {
	for i := range match {
		s = strings.ReplaceAll(s, match[i], vars[i])
	}

	return s
}

func getSyntaxHighlighting() {
	highlight, err := lang.ShellProcess.Config.Get("shell", "syntax-highlighting", types.Boolean)
	if err != nil {
		highlight = false
	}
	if highlight.(bool) {
		Prompt.SyntaxHighlighter = syntaxHighlight
	} else {
		Prompt.SyntaxHighlighter = nil
	}
}

func getHintTextEnabled() {
	showHintText, err := lang.ShellProcess.Config.Get("shell", "hint-text-enabled", types.Boolean)
	if err != nil {
		showHintText = false
	}
	if showHintText.(bool) {
		Prompt.HintText = hintText
	} else {
		Prompt.HintText = nil
	}
}

func getHintTextFormatting() {
	formatting, err := lang.ShellProcess.Config.Get("shell", "hint-text-formatting", types.String)
	if err != nil {
		formatting = ""
	}
	Prompt.HintFormatting = ansi.ExpandConsts(formatting.(string))
}

func getPreviewSettings() {
	previewImages, _ := lang.ShellProcess.Config.Get("shell", "preview-images", types.Boolean)
	Prompt.PreviewImages = previewImages.(bool)
}

var ignoreSpellCheckErr bool

func Spellchecker(r []rune) []rune {
	s := string(r)
	new, err := spellcheck.String(s)
	if err != nil && !ignoreSpellCheckErr {
		ignoreSpellCheckErr = true
		hint := fmt.Sprintf("{RED}Spellchecker error: %s{RESET} {BLUE}https://murex.rocks/docs/user-guide/spellcheck.html{RESET}", err.Error())
		Prompt.ForceHintTextUpdate(ansi.ExpandConsts(hint))
	}

	ignoreSpellCheckErr = false // reset ignore status

	return []rune(new)
}
