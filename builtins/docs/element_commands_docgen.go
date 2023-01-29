package docs

func init() {

	Definition["[["] = "# _murex_ Shell Docs\n\n## Command Reference: `[[` (element)\n\n> Outputs an element from a nested structure\n\n## Description\n\nOutputs an element from an array, map or table. Unlike **index** (`[`),\n**element** takes a path parameter which means it can work inside nested\nstructures without pipelining multiple commands together. However this\ncomes with the drawback that you can only return one element.\n\n**Element** (`[[`) also doesn't support the bang prefix (unlike) **index**.\n\nPlease note that indexes in _murex_ are counted from zero.\n\n## Usage\n\n    <stdin> -> [[ element ]] -> <stdout>\n    \n    $variable[[ element ]] -> <stdout>\n\n## Examples\n\nReturn the 2nd element in an array\n\n    » ja [0..9] -> [[ /1 ]]\n    [\n        \"1\",\n    ]\n    \nReturn the data-type and description of **config shell syntax-highlighting**\n\n    » config -> [[ /shell/syntax-highlighting/Data-Type ]]\n    bool\n\n## Detail\n\n### Element counts from zero\n\nIndexes in _murex_ behave like any other computer array in that all arrays\nstart from zero (`0`).\n\n### Alternative path separators\n\n**Element** uses the first character in the path as the separator. So the\nfollowing are all valid parameters:\n\n    » config -> [[ ,shell,syntax-highlighting,Data-Type ]]\n    bool\n    \n    » config -> [[ >shell>syntax-highlighting>Data-Type ]]\n    bool\n    \n    » config -> [[ \\|shell\\|syntax-highlighting\\|Data-Type ]]\n    bool\n    \n    » config -> [[ >shell>syntax-highlighting>Data-Type ]]\n    bool\n    \nHowever there are a few of caveats:\n\n1. Currently **element** does not support unicode separators. All separators\n   must be 1 byte characters. This limitation is highlighted as a bug, albeit\n   a low priority one. If this limitation does directly affect you then raise\n   an issue on GitHub to get the priority bumped up.\n\n2. Any shell tokens (eg pipe `|`, `;`, `}`, etc) will need to be escaped. For\n   readability reasons it is recommended not to use such characters even\n   though it is technically possible to.\n\n        # Would fail because the semi-colon is an unescaped / unquoted shell token\n        config -> [[ ;shell-syntax-highlighting;Data-Type ]]\n    \n3. Please also make sure you don't use a character that is also used inside\n   key names because keys _cannot_ be escaped. For example both of the\n   following would fail:\n\n        # Would fail because 'syntax-highlighting' and 'Data-Type' both also contain\n        # the separator character\n        config -> [[ -shell-syntax-highlighting-Data-Type ]]\n    \n        # Would fail because you cannot escape key names (escaping happens at the\n        # shell parser level rather than command parameter level)\n        config -> [[ -shell-syntax\\-highlighting-Data\\-Type ]]\n    \n### Quoting parameters\n\nIn _murex_, everything is a function. Thus even `[[` is a function name and\nthe closing `]]` is actually a last parameter. This means the recommended way\nto quote **element** parameters is to quote specific key names or the entire\npath:\n\n    » config -> [[ /shell/\"syntax-highlighting\"/Data-Type ]]\n    bool\n    \n    » config -> [[ \"|shell|syntax-highlighting|Data-Type\" ]]\n    bool\n\n## Synonyms\n\n* `[[`\n* `element`\n\n\n## See Also\n\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array"

}
