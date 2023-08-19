package docs

func init() {

	Definition["test"] = "# `test`\n\n> Murex's test framework - define tests, run tests and debug shell scripts\n\n## Description\n\n`test` is used to define tests, run tests and debug Murex shell scripts.\n\n## Usage\n\nDefine an inlined test\n\n```\ntest define test-name { json-properties }\n```\n\nDefine a state report\n\n```\ntest state name { code block }\n```\n\nDefine a unit test\n\n```\ntest unit function|private|open|event test-name { json-properties }\n```\n\nEnable or disable boolean test states (more options available in `config`)\n\n```\ntest config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]\n```\n\nDisable test mode\n\n```\n!test\n```\n\nExecute a function with testing enabled\n\n```\ntest run { code-block }\n```\n\nExecute unit test(s)\n\n```\ntest run package/module/test-name|*\n```\n\nWrite report\n\n```\ntest report\n```\n\n## Examples\n\nInlined test\n\n```\nfunction hello-world {\n    test define example {\n        \"StdoutRegex\": (^Hello World$)\n    }\n\n    out <test_example> \"Hello Earth\"\n}\n\ntest run { hello-world }\n```\n\nUnit test\n\n```\ntest unit function aliases {\n    \"PreBlock\": ({\n        alias ALIAS_UNIT_TEST=example param1 param2 param3\n    }),\n    \"StdoutRegex\": \"([- _0-9a-zA-Z]+ => .*?\\n)+\",\n    \"StdoutType\": \"str\",\n    \"PostBlock\": ({\n        !alias ALIAS_UNIT_TEST\n    })\n}\n\nfunction aliases {\n    # Output the aliases in human readable format\n    runtime --aliases -> formap name alias {\n        $name -> sprintf \"%10s => ${esccli @alias}\\n\"\n    } -> cast str\n}\n\ntest run aliases\n```\n\n## Detail\n\n### Report\n\n`test report` is only needed if `config test auto-report` is set false.\nHowever `test run` automatically enables **auto-report**.\n\nWhen the report is generated, be it automatically or manually triggered, it\nflushes the table of pending reports.\n\n## Synonyms\n\n* `test`\n* `!test`\n\n\n## See Also\n\n* [`<>` / `read-named-pipe`](../commands/namedpipe.md):\n  Reads from a Murex named pipe\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings"

}
