package docs

func init() {

	Definition["tout"] = "# _murex_ Shell Guide\n\n## Command Reference: `tout`\n\n> Print a string to the STDOUT and set it's data-type\n\n### Description\n\nWrite parameters to STDOUT without a trailing new line character. Cast the\noutput's data-type to the value of the first parameter.\n\n### Usage\n\n    tout: data-type \"string to write\" -> <stdout>\n\n### Examples\n\n    » tout: json { \"Code\": 404, \"Message\": \"Page not found\" } -> pretty\n    {\n        \"Code\": 404,\n        \"Message\": \"Page not found\"\n    }\n\n### Detail\n\n`tout` supports ANSI constants.\n\nUnlike `out`, `tout` does not append a carriage return / line feed.\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`err`](../commands/err.md):\n  Print a line to the STDERR\n* [`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* [`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [format](../commands/format.md):\n  \n* [sprintf](../commands/sprintf.md):\n  "

}
