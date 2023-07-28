package docs

func init() {

	Definition["count"] = "# `count`\n\n> Count items in a map, list or array\n\n## Description\n\n\n\n## Usage\n\n```\n<stdin> -> count: [ --duplications | --unique | --total ] -> <stdout>\n```\n\n## Examples\n\nCount number of items in a map, list or array:\n\n```\n» tout: json ([\"a\", \"b\", \"c\"]) -> count \n3\n```\n\n## Flags\n\n* `--duplications`\n    Output a JSON map of items and the number of their occurrences in a list or array\n* `--total`\n    Read an array, list or map from STDIN and output the length for that array (default behaviour)\n* `--unique`\n    Print the number of unique elements in a list or array\n* `-d`\n    Alias for `--duplications`\n* `-t`\n    Alias for `--total`\n* `-u`\n    Alias for `--unique`\n\n## Detail\n\n### Modes\n\nIf no flags are set, `count` will default to using `--total`.\n\n#### Total: `--total` / `-t`\n\nThis will read an array, list or map from STDIN and output the length for\nthat array.\n\n```\n» a [25-Dec-2020..05-Jan-2021] -> count \n12\n```\n\n> This also replaces the older `len` method.\n\nPlease note that this returns the length of the _array_ rather than string.\nFor example `out \"foobar\" -> count` would return `1` because an array in the\n`str` data type would be new line separated (eg `out \"foo\\nbar\" -> count`\nwould return `2`). If you need to count characters in a string and are\nrunning POSIX (eg Linux / BSD / OSX) then it is recommended to use `wc`\ninstead. But be mindful that `wc` will also count new line characters.\n\n```\n» out: \"foobar\" -> count\n1\n\n» out: \"foo\\nbar\" -> count\n2\n\n» out: \"foobar\" -> wc: -c\n7\n\n» out: \"foo\\nbar\" -> wc: -c\n8\n\n» printf: \"foobar\" -> wc: -c\n6\n# (printf does not print a trailing new line)\n```\n\n#### Duplications: `--duplications` / `-d`\n\nThis returns a JSON map of items and the number of their occurrences in a list\nor array.\n\nFor example in the quote below, only the word \"the\" is repeated so that entry\nwill have a value of `2` while ever other entry has a value of `1` because they\nonly appear once in the quote.\n\n```\n» out: \"the quick brown fox jumped over the lazy dog\" -> jsplit: \\s -> count: --duplications\n{\n    \"brown\": 1,\n    \"dog\": 1,\n    \"fox\": 1,\n    \"jumped\": 1,\n    \"lazy\": 1,\n    \"over\": 1,\n    \"quick\": 1,\n    \"the\": 2\n}\n```\n\n#### Unique: `--unique` / `-u`\n\nReturns the number of unique elements in a list or array.\n\nFor example in the quote below, only the word \"the\" is repeated, thus the\nunique count should be one less than the total count:\n\n```\n» out \"the quick brown fox jumped over the lazy dog\" -> jsplit \\s -> count --unique\n8\n» out \"the quick brown fox jumped over the lazy dog\" -> jsplit \\s -> count --total\n9\n```\n\n## Synonyms\n\n* `count`\n* `len`\n\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [`ta` (mkarray)](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table"

}
