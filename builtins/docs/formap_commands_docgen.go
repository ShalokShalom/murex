package docs

func init() {

	Definition["formap"] = "# `formap` - Command Reference\n\n> Iterate through a map or other collection of data\n\n## Description\n\n`formap` is a generic tool for iterating through a map, table or other\nsequences of data similarly like a `foreach`. In fact `formap` can even be\nused on array too.\n\nUnlike `foreach`, `formap`'s default output is `str`, so each new line will be\ntreated as a list item. This behaviour will differ if any additional flags are\nused with `foreach`, such as `--jmap`.\n\n## Usage\n\n`formap` writes a list:\n\n```\n<stdin> -> foreach variable { code-block } -> <stdout>\n```\n\n`formap` writes to a buffered JSON map:\n\n```\n<stdin> -> formap --jmap key value { code-block (map key) } { code-block (map value) } -> <stdout>\n```\n\n## Examples\n\nFirst of all lets assume the following dataset:\n\n```\nset json people={\n    \"Tom\": {\n        \"Age\": 32,\n        \"Gender\": \"Male\"\n    },\n    \"Dick\": {\n        \"Age\": 43,\n        \"Gender\": \"Male\"\n    },\n    \"Sally\": {\n        \"Age\": 54,\n        \"Gender\": \"Female\"\n    }\n}\n```\n\nWe can create human output from this:\n\n```\n» $people -> formap key value { out \"$key is $value[Age] years old\" }\nSally is 54 years old\nTom is 32 years old\nDick is 43 years old\n```\n\n> Please note that maps are intentionally unsorted so you cannot guarantee the\n> order of the output produced even if the input has been superficially set in\n> a specific order.\n\nWith `--jmap` we can turn that structure into a new structure:\n\n```\n» $people -> formap --jmap key value { $key } { $value[Age] }\n{\n    \"Dick\": \"43\",\n    \"Sally\": \"54\",\n    \"Tom\": \"32\"\n} \n```\n\n## Flags\n\n* `--jmap`\n    Write a `json` map to STDOUT instead of an array\n\n## Detail\n\n`formap` can also work against arrays and tables as well. However `foreach` is\na much better tool for ordered lists and tables can look a little funky when\nwhen there are more than 2 columns. In those instances you're better off using\n`[` (index) to specify columns and then `tabulate` for any data transformation.\n\n### Meta values\n\nMeta values are a JSON object stored as the variable `$.`. The meta variable\nwill get overwritten by any other block which invokes meta values. So if you\nwish to persist meta values across blocks you will need to reassign `$.`, eg\n\n```\n%[1..3] -> foreach {\n    meta_parent = $.\n    %[7..9] -> foreach {\n        out \"$(meta_parent.i): $.i\"\n    }\n}\n```\n\nThe following meta values are defined:\n\n* `i`: iteration number\n\n## See Also\n\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`for`](../commands/for.md):\n  A more familiar iteration loop to existing developers\n* [`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [`json` ](../types/json.md):\n  JavaScript Object Notation (JSON)\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`tabulate`](../commands/tabulate.md):\n  Table transformation tools\n* [`while`](../commands/while.md):\n  Loop until condition false\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table"

}
