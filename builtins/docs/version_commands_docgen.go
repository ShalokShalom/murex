package docs

func init() {

	Definition["version"] = "# _murex_ Shell Docs\n\n## Command Reference: `version` \n\n> Get _murex_ version\n\n## Description\n\nReturns _murex_ version number\n\n## Usage\n\n    version [ flags ] -> <stdout>\n\n## Examples\n\nRan without any parameters\n\n    » version\n    murex: 0.51.1200 BETA\n    \nRan with the `--no-app-name` parameter\n\n    » version --no-app-name\n    0.51.1200 BETA\n    \nRan with the `--short` parameter\n\n    » version --short\n    0.51\n\n## Flags\n\n* `--no-app-name`\n    Returns full version string minus app name\n* `--short`\n    Returns only the major and minor version as a `num` data-type\n\n## See Also\n\n* [commands/`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [commands/`function`](../commands/function.md):\n  Define a function block\n* [commands/`murex-parser` ](../commands/murex-parser.md):\n  Runs the _murex_ parser against a block of code \n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [commands/`source` ](../commands/source.md):\n  Import _murex_ code from another file of code block\n* [commands/args](../commands/args.md):\n  \n* [commands/params](../commands/params.md):\n  "

}