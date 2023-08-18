package docs

func init() {

	Definition["man-get-flags"] = "# `man-get-flags` \n\n> Parses man page files for command line flags \n\n## Description\n\nSometimes you might want to programmatically search `man` pages for any\nsupported flag. Particularly if you're writing a dynamic autocompletion.\n`man-get-flags` does this and returns a JSON document.\n\nYou can either pipe a man page to `man-get-flags`, or pass the name of the\ncommand as a parameter.\n\n`man-get-flags` returns a JSON document. Either an array or an object,\ndepending on what flags (if any) are passed.\n\nIf no flags are passed, `man-get-flags` will default to just parsing the man\npage for anything that looks like a flag (ie no descriptions or other detail).\n\n## Usage\n\n```\n<stdin> -> man-get-flags [--descriptions] -> <stdout>\n\nman-get-flags command [--descriptions] -> <stdout>\n```\n\n## Examples\n\n```\n» man-get-flags --descriptions find -> [{$.key =~ 'regex'}]\n{\n    \"-iregex\": \"eg: pattern -- Like -regex, but the match is case insensitive.\",\n    \"-regex\": \"eg: pattern -- True if the whole path of the file matches pattern using regular expression.  To match a file named “./foo/xyzzy”, you can use the regular expression “.*/[xyz]*” or “.*/foo/.*”, but not “xyzzy” or “/foo/”.\"\n}\n```\n\n## Flags\n\n* `--descriptions`\n    return a map of flags with their described usage\n* `-d`\n    shorthand for `--descriptions`\n\n## Detail\n\n### Limitations\n\nDue to the freeform nature of man pages - that they're intended to be human\nreadable rather than machine readable - and the flexibility that developers\nhave to parse command line parameters however they wish, there will always be\na margin for error with how reliably any parser can autodetect parameters. one\nrequirement is that flags are hyphen prefixed, eg `--flag`.\n\n## See Also\n\n* [`man-summary`](../commands/man-summary.md):\n  Outputs a man page summary of a command\n* [`murex-docs`](../commands/murex-docs.md):\n  Displays the man pages for Murex builtins\n* [`summary` ](../commands/summary.md):\n  Defines a summary help text for a command"

}
