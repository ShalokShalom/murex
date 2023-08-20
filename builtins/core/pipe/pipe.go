package cmdpipe

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("pipe", cmdPipe, types.Null)
	lang.DefineFunction("!pipe", cmdClosePipe, types.Null)
}

func cmdPipe(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("missing parameters")
	}

	// import the registered pipes
	supportedFlags := make(map[string]string)
	pipes := stdio.DumpPipes()
	for i := range pipes {
		supportedFlags["--"+pipes[i]] = types.String
	}

	// define cli flags
	flags, additional, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: true,
		Flags:           supportedFlags,
	})

	if err != nil {
		return err
	}

	if len(additional) == 0 {
		return errors.New("no name specified for named pipe. Usage: `pipe name [ --pipe-type creation-data ]")
	}

	if len(flags) > 1 {
		return errors.New("too many types of pipe specified. Please use only one flag per")
	}

	for flag := range flags {
		return lang.GlobalPipes.CreatePipe(additional[0], flag[2:], flags[flag])
	}

	for _, name := range additional {
		err := lang.GlobalPipes.CreatePipe(name, "std", "")
		if err != nil {
			return err
		}
	}

	return nil
}

func cmdClosePipe(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	var names []string

	if p.IsMethod {
		p.Stdin.ReadArray(p.Context, func(b []byte) {
			names = append(names, string(b))
		})

		if len(names) == 0 {
			return errors.New("stdin contained a zero length array")
		}

	} else {
		if p.Parameters.Len() == 0 {
			return errors.New("no pipes listed for closing")
		}

		names = p.Parameters.StringArray()
	}

	for _, name := range names {
		if err := lang.GlobalPipes.Close(name); err != nil {
			return err
		}
	}

	return nil
}
