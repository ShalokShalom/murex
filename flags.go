package main

import (
	"flag"
	"fmt"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/lang/types"
)

var (
	fCommand     string
	fInteractive bool
	fSource      []string
	fLoadMods    bool
	fEcho        bool
	fHelp1       bool
	fHelp2       bool
	fVersion1    bool
	fVersion2    bool
	fSh          bool
	fRunTests    bool
)

func readFlags() {
	flag.StringVar(&fCommand, "c", "", "Run code block - read from parameters")
	flag.BoolVar(&fInteractive, "i", false, "Start interactive shell after -c execution")
	flag.BoolVar(&fLoadMods, "load-modules", false, "Load modules and profile when in non-interactive mode ")

	flag.BoolVar(&fHelp1, "h", false, "Help")
	flag.BoolVar(&fHelp2, "help", false, "Help")

	flag.BoolVar(&fVersion1, "v", false, "Version")
	flag.BoolVar(&fVersion2, "version", false, "Version")

	flag.BoolVar(&debug.Enabled, "debug", false, "Debug mode (for debugging murex code. This can also be enabled from inside the shell.")
	flag.BoolVar(&fRunTests, "run-tests", false, "Run all tests and exit")
	flag.BoolVar(&fEcho, "echo", false, "Echo on")
	flag.BoolVar(&fSh, "murex", false, "")

	flag.BoolVar(&lang.FlagTry, "try", false, "Enable a global `try` block")
	flag.BoolVar(&lang.FlagTryPipe, "trypipe", false, "Enable a global `trypipe` block")

	flag.Parse()

	if fHelp1 || fHelp2 {
		fmt.Fprintf(tty.Stdout, "%s v%s\n", app.Name, app.Version())
		flag.Usage()
		lang.Exit(1)
	}

	if fVersion1 || fVersion2 {
		fmt.Fprintf(tty.Stdout, "%s v%s\n", app.Name, app.Version())
		fmt.Fprintf(tty.Stdout, "%s\n%s\n", app.License, app.Copyright)
		lang.Exit(0)
	}

	config.InitConf.Define("proc", "echo", config.Properties{
		Description: "Echo shell functions",
		Default:     fEcho,
		DataType:    types.Boolean,
	})

	fSource = flag.Args()
}
