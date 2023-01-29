package docs

func init() {

	Definition["base64"] = "# _murex_ Shell Docs\n\n## Optional Command Reference: `base64` \n\n> Encode or decode a base64 string\n\n## Description\n\nAn optional builtin to encode or decode a base64 string.\n\n## Usage\n\n    <stdin> -> base64 -> <stdout>\n    \n    <stdin> -> !base64 -> <stdout>\n\n## Examples\n\nEncode base64 string\n\n    » out: \"Hello, World!\" -> base64\n    SGVsbG8sIFdvcmxkIQo=\n    \nDecode base64 string\n\n    » out: \"SGVsbG8sIFdvcmxkIQo=\" -> !base64\n    Hello, World!\n\n## Detail\n\n`base64` is very simplistic - particularly when compared to its GNU coreutil\n(for example) counterpart. If you want to use the `base64` binary on Linux\nor similar platforms then you will need to launch with the `exec` builtin:\n\n    » out: \"Hello, World!\" -> exec: base64\n    SGVsbG8sIFdvcmxkIQo=\n    \n    » out: \"SGVsbG8sIFdvcmxkIQo=\" -> exec: base64 -d\n    Hello, World!\n    \nHowever for simple tasks this builtin will out perform external tools because\nit doesn't require the OS fork processes.\n\n## Synonyms\n\n* `base64`\n* `!base64`\n\n\n## See Also\n\n* [`!bz2` ](../optional/bz2.md):\n  Decompress a bz2 file\n* [`escape`](../commands/escape.md):\n  Escape or unescape input \n* [`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL\n* [`gz` ](../optional/gz.md):\n  Compress or decompress a gzip file"

}
