package docs

func init() {

	Definition["pt"] = "# `pt`\n\n> Pipe telemetry. Writes data-types and bytes written\n\n## Description\n\nPipe telemetry, `pt`, writes statistics about the pipeline. The telemetry is written\ndirectly to the OS's STDERR so to preserved the pipeline.\n\n## Usage\n\n```\n<stdin> -> pt -> <stdout>\n```\n\n## Examples\n\n```\ncurl -s https://example.com/bigfile.bin -> pt -> > bigfile.bin\n```\n\n(though Murex does also have it's own HTTP clients, `get`, `post` and\n`getfile`)\n\n## See Also\n\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [`get`](../commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [`getfile`](../commands/getfile.md):\n  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.\n* [`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return"

}
