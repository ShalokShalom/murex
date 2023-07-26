package docs

func init() {

	Definition["history"] = "# `history` - Command Reference\n\n> Outputs murex's command history\n\n## Description\n\nOutputs _mutex_'s command history.\n\n## Usage\n\n```\nhistory -> <stdout>\n```\n\n## Examples\n\n```\n» history\n...\n{\n    \"Index\": 16782,\n    \"DateTime\": \"2019-01-19T22:43:21.124273664Z\",\n    \"Block\": \"tout: json ([\\\"a\\\", \\\"b\\\", \\\"c\\\"]) -\\u003e len\"\n},\n{\n    \"Index\": 16783,\n    \"DateTime\": \"2019-01-19T22:50:42.114986768Z\",\n    \"Block\": \"clear\"\n},\n{\n    \"Index\": 16784,\n    \"DateTime\": \"2019-01-19T22:51:39.82077789Z\",\n    \"Block\": \"map { tout: json ([\\\"key 1\\\", \\\"key 2\\\", \\\"key 3\\\"]) }\"\n},\n...\n```\n\n## Detail\n\nThe history file is typically located on disk in a file called `~/.murex.history`.\n\n## See Also\n\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex"

}
