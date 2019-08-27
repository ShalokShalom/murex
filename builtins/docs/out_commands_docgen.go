package docs

func init() {

	Definition["out"] = "# _murex_ Shell Guide\n\n## Command Reference: `out`\n\n> `echo` a string to the STDOUT with a trailing new line character\n\n### Description\n\nWrite parameters to STDOUT with a trailing new line character.\n\n### Usage\n\n    out: string to write -> <stdout>\n\n### Examples\n\n    » out Hello, World!\n    Hello, World!\n    \nFor compatibility with other shells, `echo` is also supported:\n\n    » echo Hello, World!\n    Hello, World!\n\n### Detail\n\n`out` / `echo` output as `string` data-type. This can be changed by casting\n(`cast`) or using the `tout` function.\n\n#### ANSI Constants\n\n`out` supports ANSI constants.\n\n### Synonyms\n\n* `out`\n* `echo`\n\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`err`](../commands/err.md):\n  Print a line to the STDERR\n* [`pt`](../commands/pt.md):\n  Pipe telemetry. Writes data-types and bytes written\n* [`read`](../commands/read.md):\n  `read` a line of input from the user and store as a variable\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [`tread`](../commands/tread.md):\n  `read` a line of input from the user and store as a user defined *typed* variable\n* [sprintf](../commands/sprintf.md):\n  "

}
