package docs

func init() {

	Definition["openagent"] = "# `openagent` - Command Reference\n\n> Creates a handler function for `open`\n\n## Description\n\n`openagent` creates and destroys handler functions for writing data to the\nterminal when accessed via `open` and STDOUT is a TTY.\n\n## Usage\n\nDisplay code block for an associated data-type:\n\n```\nopenagent get data-type\n```\n\nDefine an `open` handler function:\n\n```\nopenagent set data-type { code-block }\n```\n\nUndefine an `open` handler:\n\n```\n!openagent data-type\n```\n\n## Detail\n\n### FileRef\n\nIt is possible to track which shell script or module installed what `open`\nhandler by checking `runtime --open-agents` and checking it's **FileRef**.\n\n## Synonyms\n\n* `openagent`\n* `!openagent`\n\n\n## See Also\n\n* [FileRef](../user-guide/fileref.md):\n  How to track what code was loaded and from where\n* [Modules and Packages](../user-guide/modules.md):\n  An introduction to Murex modules and packages\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`open`](../commands/open.md):\n  Open a file with a preferred handler\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex"

}
