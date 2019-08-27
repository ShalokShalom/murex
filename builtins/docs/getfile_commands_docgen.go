package docs

func init() {

	Definition["getfile"] = "# _murex_ Shell Guide\n\n## Command Reference: `getfile`\n\n> Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.\n\n### Description\n\nFetches a resource from a URL - setting STDOUT data-type\n\n### Usage\n\n    getfile url -> <stdout>\n\n### Examples\n\n    getfile google.com \n\n### Detail\n\nThis simply fetches a resource (via HTTP GET request) from a URL and returns the\nbyte stream to STDOUT. It will set STDOUT's data-type based on MIME defined in\nthe `Content-Type` HTTP header.\n\nIt is recommended that you only use this command if you're pipelining the output\n(eg writing to file or passing on to another function). If you just want to\nrender the output to the terminal then use `open` which has hooks for smart\nterminal rendering.\n\n#### Configurable options\n\n`getfile` has a number of behavioral options which can be configured via\n_murex_'s standard `config` tool:\n\n    config: -> [ http ]\n    \nTo change a default, for example the user agent string:\n\n    config: set http user-agent \"bob\"\n    getfile: google.com\n    \nThis enables sane, repeatable and readable defaults. Read the documents on\n`config` for more details about it's usage and the rational behind the command.\n\n### See Also\n\n* [`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [`get`](../commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return\n* [open](../commands/open.md):\n  "

}
