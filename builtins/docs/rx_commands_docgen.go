package docs

func init() {

	Definition["rx"] = "# `rx`\n\n> Regexp pattern matching for file system objects (eg `.*\\\\.txt`)\n\n## Description\n\nReturns a list of files and directories that match a regexp pattern.\n\nOutput is a JSON list.\n\n## Usage\n\n```\nrx: pattern -> <stdout>\n\n!rx: pattern -> <stdout>\n\n<stdin> -> rx: pattern -> <stdout>\n\n<stdin> -> !rx: pattern -> <stdout>\n```\n\n## Examples\n\nInline regex file matching:\n\n```\ncat: @{ rx: '.*\\.txt' }\n```\n\nWriting a list of files to disk:\n\n```\nrx: '.*\\.go' |> filelist.txt\n```\n\nChecking if files exist:\n\n```\nif { rx: somefiles.* } then {\n    # files exist\n}\n```\n\nChecking if files do not exist:\n\n```\n!if { rx: somefiles.* } then {\n    # files do not exist\n}\n```\n\nReturn all files apart from text files:\n\n```\n!g: '\\.txt$'\n```\n\nFiltering a file list based on regexp matches file:\n\n```\nf: +f -> rx: '.*\\.txt'\n```\n\nRemove any regexp file matches from a file list:\n\n```\nf: +f -> !rx: '.*\\.txt'\n```\n\n## Detail\n\n### Traversing Directories\n\nUnlike globbing (`g`) which can traverse directories (eg `g: /path/*`), `rx` is\nonly designed to match file system objects in the current working directory.\n\n`rx` uses Go (lang)'s standard regexp engine.\n\n### Inverse Matches\n\nIf you want to exclude any matches based on wildcards, rather than include\nthem, then you can use the bang prefix. eg\n\n```\n» rx: READ*                                                                                                                                                              \n[\n    \"README.md\"\n]\n\nmurex-dev» !rx: .*\nError in `!rx` (1,1): No data returned.\n```\n\n### When Used As A Method\n\n`!rx` first looks for files that match its pattern, then it reads the file list\nfrom STDIN. If STDIN contains contents that are not files then `!rx` might not\nhandle those list items correctly. This shouldn't be an issue with `rx` in its\nnormal mode because it is only looking for matches however when used as `!rx`\nany items that are not files will leak through.\n\nThis is its designed feature and not a bug. If you wish to remove anything that\nalso isn't a file then you should first pipe into either `g: *`, `rx: .*`, or\n`f +f` and then pipe that into `!rx`.\n\nThe reason for this behavior is to separate this from `!regexp` and `!match`.\n\n## Synonyms\n\n* `rx`\n* `!rx`\n\n\n## See Also\n\n* [`f`](../commands/f.md):\n  Lists or filters file system objects (eg files)\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [`match`](../commands/match.md):\n  Match an exact value in an array\n* [`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings"

}
