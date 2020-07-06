package docs

func init() {

	Definition["while"] = "# _murex_ Shell Docs\n\n## Command Reference: `while`\n\n> Loop until condition false\n\n## Description\n\n`while` loops until loops until **condition** is false.\n\nNormally the **conditional** and executed code block are 2 separate parameters\nhowever you can call `while` with just 1 parameter where the code block acts\nas both the conditional and the code to be ran.\n\n## Usage\n\nUntil true\n\n    while { condition } { code-block } -> <stdout>\n    \n    while { code-block } -> <stdout>\n    \nUntil false\n\n    !while { condition } { code-block } -> <stdout>\n    \n``\n!while { code-block } -> <std\n\n## Examples\n\n`while` **$i** is less then **5**\n\n    » let i=0; while { =i<5 } { let i=i+1; out $i }\n    1\n    2\n    3\n    4\n    5\n    \n    » let i=0; while { let i=i+1; = i<5; out }\n    true\n    true\n    true\n    true\n    false\n    \n`while` **$i** is _NOT_ greater than or equal to **5**\n\n    » let i=0; !while { =i>=5 } { let i=i+1; out $i }\n    1\n    2\n    3\n    4\n    5\n    \n    » let i=0; while { let i=i+1; = i>=5; out }\n    true\n    true\n    true\n    true\n    false\n\n## Synonyms\n\n* `while`\n* `!while`\n\n\n## See Also\n\n* [commands/`err`](../commands/err.md):\n  Print a line to the STDERR\n* [commands/`for`](../commands/for.md):\n  A more familiar iteration loop to existing developers\n* [commands/`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable\n* [commands/`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [commands/formap](../commands/formap.md):\n  "

}
