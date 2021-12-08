package docs

func init() {

	Definition["prepend"] = "# _murex_ Shell Docs\n\n## Command Reference: `prepend` \n\n> Add data to the start of an array\n\n## Description\n\n`prepend` a data to the start of an array.\n\n## Usage\n\n    <stdin> -> prepend: value -> <stdout>\n\n## Examples\n\n    » a: [January..December] -> prepend: 'New Year'\n    New Year\n    January\n    February\n    March\n    April\n    May\n    June\n    July\n    August\n    September\n    October\n    November\n    December\n\n## Detail\n\nIt's worth noting that `prepend` and `append` are not data type aware. So \nany integers in data type aware structures will be converted into strings:\n\n    » tout: json [1,2,3] -> prepend: new \n    [\n        \"new\",\n        \"1\",\n        \"2\",\n        \"3\"\n    ]\n\n## See Also\n\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`len` ](../commands/len.md):\n  Outputs the length of an array\n* [commands/`match`](../commands/match.md):\n  Match an exact value in an array\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [commands/`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings"

}
