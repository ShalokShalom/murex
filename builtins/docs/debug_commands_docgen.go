package docs

func init() {

	Definition["debug"] = "# _murex_ Shell Docs\n\n## Command Reference: `debug`\n\n> Debugging information\n\n## Description\n\n`debug` has two modes: as a function and as a method.\n\n### Debug Method\n\nThis usage will return debug information about the previous function ran.\n\n### Debug Function:\n\nThis will enable or disable debugging mode.\n\n## Usage\n\n    <stdin> -> debug -> <stdout>\n    \n    debug: boolean -> <stdout>\n    \n    debug -> <stdout>\n\n## Examples\n\nReturn debugging information on the previous function:\n\n    » echo: \"hello, world!\" -> debug \n    {\n        \"DataType\": {\n            \"Go\": \"[]string\",\n            \"Murex\": \"str\"\n        },\n        \"Process\": {\n            \"Context\": {\n                \"Context\": 0\n            },\n            \"Stdin\": {},\n            \"Stdout\": {},\n            \"Stderr\": {},\n            \"Parameters\": {\n                \"Params\": [\n                    \"hello, world!\"\n                ],\n                \"Tokens\": [\n                    [\n                        {\n                            \"Type\": 0,\n                            \"Key\": \"\"\n                        }\n                    ],\n                    [\n                        {\n                            \"Type\": 1,\n                            \"Key\": \"hello, world!\"\n                        }\n                    ],\n                    [\n                        {\n                            \"Type\": 0,\n                            \"Key\": \"\"\n                        }\n                    ]\n                ]\n            },\n            \"ExitNum\": 0,\n            \"Name\": \"echo\",\n            \"Id\": 3750,\n            \"Exec\": {\n                \"Pid\": 0,\n                \"Cmd\": null,\n                \"PipeR\": null,\n                \"PipeW\": null\n            },\n            \"PromptGoProc\": 1,\n            \"Path\": \"\",\n            \"IsMethod\": false,\n            \"IsNot\": false,\n            \"NamedPipeOut\": \"out\",\n            \"NamedPipeErr\": \"err\",\n            \"NamedPipeTest\": \"\",\n            \"State\": 7,\n            \"IsBackground\": false,\n            \"LineNumber\": 1,\n            \"ColNumber\": 1,\n            \"RunMode\": 0,\n            \"Config\": {},\n            \"Tests\": {\n                \"Results\": null\n            },\n            \"Variables\": {},\n            \"FidTree\": [\n                0,\n                3750\n            ],\n            \"CreationTime\": \"2019-01-20T00:00:52.167127131Z\",\n            \"StartTime\": \"2019-01-20T00:00:52.167776212Z\"\n        }\n    }\n    \nEnable or disable debug mode:\n\n    » debug: on\n    true\n    \n    » debug: off\n    false\n    \nOutput whether debug mode is enabled or disabled:\n\n    » debug\n    false\n\n## Detail\n\nWhen enabling or disabling debug mode, because the parameter is a murex\nboolean type, it means you can use other boolean terms. eg\n\n    # enable debugging\n    » debug 1\n    » debug on\n    » debug yes\n    » debug true\n    \n    # disable debugging\n    » debug 0\n    » debug off\n    » debug no\n    » debug false\n    \nIt is also worth noting that the debugging information needs to be written\ninto the Go source code rather than in _murex_'s shell scripting language.\nIf you require debugging those processes then please use _murex_'s `test`\nframework\n\n## See Also\n\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [`test`](../commands/test.md):\n  _murex_'s test framework - define tests, run tests and debug shell scripts"

}
