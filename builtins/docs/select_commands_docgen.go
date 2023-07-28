package docs

func init() {

	Definition["select"] = "# `select` \n\n> Inlining SQL into shell pipelines\n\n## Description\n\n`select` imports tabulated data into an in memory sqlite3 database and\nexecutes SQL queries against the data. It returns a table of the same\ndata type as the input type\n\n## Usage\n\n```\n<stdin> -> select * | ... WHERE ... -> <stdout>\n\nselect * | ... FROM file[.gz] WHERE ... -> <stdout>\n```\n\n## Examples\n\nList a count of all the processes running against each user ID:\n\n```\n» ps aux -> select count(*), user GROUP BY user ORDER BY 1\ncount(*) USER\n1       _analyticsd\n1       _applepay\n1       _atsserver\n1       _captiveagent\n1       _cmiodalassistants\n1       _ctkd\n1       _datadetectors\n1       _displaypolicyd\n1       _distnote\n1       _gamecontrollerd\n1       _hidd\n1       _iconservices\n1       _installcoordinationd\n1       _mdnsresponder\n1       _netbios\n1       _networkd\n1       _reportmemoryexception\n1       _timed\n1       _usbmuxd\n2       _appleevents\n3       _assetcache\n3       _fpsd\n3       _nsurlsessiond\n3       _softwareupdate\n4       _windowserver\n5       _coreaudiod\n6       _spotlight\n7       _locationd\n144     root\n308     foobar\n```\n\n```\n\nselect count(*)\n```\n\n## Detail\n\n### Default Table Name\n\nThe table created is called `main`, however you do not need to include a `FROM`\ncondition in your SQL as Murex will inject `FROM main` into your SQL if it is\nmissing. In fact, it is recommended that you exclude `FROM` from your SQL\nqueries for the sake of brevity.\n\n### `config` Options\n\n`select`'s behavior is configurable:\n\n```\n» config -> [ select ]\n{\n    \"fail-irregular-columns\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": false,\n        \"Description\": \"When importing a table into sqlite3, fail if there is an irregular number of columns\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": false\n    },\n    \"merge-trailing-columns\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"When importing a table into sqlite3, if `fail-irregular-columns` is set to `false` and there are more columns than headings, then any additional columns are concatenated into the last column (space delimitated). If `merge-trailing-columns` is set to `false` then any trailing columns are ignored\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": true\n    },\n    \"print-headings\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"Print headings when writing results\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": true\n    },\n    \"table-includes-headings\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"When importing a table into sqlite3, treat the first row as headings (if `false`, headings are Excel style column references starting at `A`)\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": true\n    }\n}\n```\n\n(See below for how to use `config`)\n\n### Read All vs Sequential Reads\n\nAt present, `select` only supports reading the entire table from STDIN before\nimporting that data into sqlite3. There is some prototype code being written to\nsupport sequential imports but this is hugely experimental and not yet enabled.\n\nThis might make `select` unsuitable for large datasets.\n\n### Early Release\n\nThis is a very early release so there almost certainly will be bugs hiding.\nWhich is another reason why this is currently only an optional builtin.\n\nIf you do run into any issues then please raise them on [Github](https://github.com/lmorg/murex/issues).\n\n## Synonyms\n\n* `select`\n\n\n## See Also\n\n* [`*` (generic) ](../types/generic.md):\n  generic (primitive)\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`csv` ](../types/csv.md):\n  CSV files (and other character delimited tables)\n* [`jsonl` ](../types/jsonl.md):\n  JSON Lines\n* [v2.1](../changelog/v2.1.md):\n  This release comes with support for inlining SQL and some major bug fixes plus a breaking change for `config`. Please read for details."

}
