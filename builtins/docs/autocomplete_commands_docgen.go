package docs

func init() {

	Definition["autocomplete"] = "# `autocomplete` - Command Reference\n\n> Set definitions for tab-completion in the command line\n\n## Description\n\n`autocomplete` digests a JSON schema and uses that to define the tab-\ncompletion rules for suggestions in the interactive command line.\n\n## Usage\n\n    autocomplete get [ command ] -> <stdout>\n    \n    autocomplete set command { mxjson }\n\n## Flags\n\n* `get`\n    output all autocompletion schemas\n* `set`\n    define a new autocompletion schema\n\n## Detail\n\n### Undefining autocomplete\n\nCurrently there is no support for undefining an autocompletion rule however\nyou can overwrite existing rules.\n\n## Directives\n\nThe directives are listed below. Headings are formatted as follows:\n\n    \"DirectiveName\": json data-type (default value)\n    \nWhere \"default value\" is what will be auto-populated at run time if you don't\ndefine an autocomplete schema manually. **zls** stands for zero-length string\n(ie: \"\").\n\n<div id=\"toc\">\n\n- [\"Alias\": string (zls)](#alias-string-zls)\n- [\"AllowAny\": boolean (false)](#allowany-boolean-false)\n- [\"AllowMultiple\": boolean (false)](#allowmultiple-boolean-false)\n- [\"AnyValue\": boolean (false)](#anyvalue-boolean-false)\n- [\"AutoBranch\": boolean (false)](#autobranch-boolean-false)\n- [\"CacheTTL\": int (5)](#cachettl-int-5)\n- [\"Dynamic\": string (zls)](#dynamic-string-zls)\n- [\"DynamicDesc\": string (zls)](#dynamicdesc-string-zls)\n- [\"ExecCmdline\": boolean (false)](#execcmdline-boolean-false)\n- [\"FileRegexp\": string (zls)](#fileregexp-string-zls)\n- [\"FlagValues\": map of arrays (null)](#flagvalues-map-of-arrays-null)\n  - [Defaults for matched flags](#defaults-for-matched-flags)\n  - [Defaults for any flags (including unmatched)](#defaults-for-any-flags-including-unmatched)\n- [\"Flags\": array of strings (auto-populated from man pages)](#flags-array-of-strings-auto-populated-from-man-pages)\n- [\"FlagsDesc\": map of strings (null)](#flagsdesc-map-of-strings-null)\n- [\"Goto\": string (zls)](#goto-string-zls)\n- [\"IgnorePrefix\": boolean (false)](#ignoreprefix-boolean-false)\n- [\"IncDirs\": boolean (false)](#incdirs-boolean-false)\n- [\"IncExeAll\": boolean (false)](#incexeall-boolean-false)\n- [\"IncExePath\": boolean (false)](#incexepath-boolean-false)\n- [\"IncFiles\": boolean (true)](#incfiles-boolean-true)\n- [\"IncManPage\": boolean (false)](#incmanpage-boolean-false)\n- [\"ListView\": boolean (false)](#listview-boolean-false)\n- [\"NestedCommand\": boolean (false)](#nestedcommand-boolean-false)\n- [\"Optional\": boolean (false)](#optional-boolean-false)\n\n\n</div>\n\n### \"Alias\": string (zls)\n\nAliases are used inside **FlagValues** as a way of pointing one flag to another\nwithout duplicating code. eg `-v` and `--version` might be the same flag. Or\n`-?`, `-h` and `--help`. With **Alias** you can write the definitions for one\nflag and then point all the synonyms as an alias to that definition.\n\n### \"AllowAny\": boolean (false)\n\nThe way autocompletion works in Murex is the suggestion engine looks for\nmatches and if it fines one, it then moves onto the next index in the JSON\nschema. This means unexpected values typed in the interactive terminal will\nbreak the suggestion engine's ability to predict what the next expected\nparameter should be. Setting **AllowAny** to `true` tells the suggestion\nengine to accept any value as the next parameter thus allowing it to then\npredict the next parameter afterwards.\n\nThis directive isn't usually necessary because such fields are often the last\nparameter or most parameters can be detectable with a reasonable amount of\neffort. However **AllowAny** is often required for more complex command line\ntools.\n\n### \"AllowMultiple\": boolean (false)\n\nSet to `true` to enable multiple parameters following the same rules as defined\nin this index. For example the following will suggest directories on each tab\nfor multiple parameters:\n\n    autocomplete set example { [{\n        \"IncDirs\": true,\n        \"AllowMultiple\": true\n    }] }\n    \n### \"AnyValue\": boolean (false)\n\nDeprecated. Please use **AllowAny** instead.\n\n### \"AutoBranch\": boolean (false)\n\nUse this in conjunction with **Dynamic**. If the return is an array of paths,\nfor example `[ \"/home/foo\", \"/home/bar\" ]` then **AutoBranch** will return\nthe following patterns in the command line:\n\n    » example [tab]\n    # suggests \"/home/\"\n    \n    » example /home/[tab]\n    # suggests \"/home/foo\" and \"/home/bar\"\n    \nPlease note that **AutoBranch**'s behavior is also dependant on a \"shell\"\n`config` setting, recursive-enabled\":\n\n    » config get shell recursive-enabled\n    true\n    \n### \"CacheTTL\": int (5)\n\nDynamic autocompletions (via **Dynamic** or **DynamicDesc**) are cached to\nimprove interactivity performance. By default the cache is very small but you\ncan increase that cache or even disable it entirely. Setting this value will\ndefine the duration (in seconds) to cache that autocompletion.\n\nIf you wish to disable this then set **CacheTTL** to `-1`.\n\nThis directive needs to live in the very first definition and affects all\nautocompletes within the rest of the command. For example\n\n    autocomplete set foobar { [\n        {\n            \"Flags\": [ \"--foo\", \"--bar\" ],\n            \"CacheTTL\": 60\n        },\n        {\n            \"Dynamic\": ({\n                a: [Monday..Friday]\n                sleep: 3\n            })\n        }\n    ] }\n    \nHere the days of the week take 3 seconds to show up as autocompletion\nsuggestions the first time and instantly for the next 60 seconds after.\n\n### \"Dynamic\": string (zls)\n\nThis is a Murex block which returns an array of suggestions.\n\nCode inside that block are executed like a function and the parameters will\nmirror the same as those parameters entered in the interactive terminal.\n\nTwo variables are created for each **Dynamic** function:\n\n* `ISMETHOD`: `true` if the command being autocompleted is going to run as a\n              pipelined method. `false` if it isn't.\n\n* `PREFIX`: contains the partial term. For example if you typed `hello wor[tab]`\n            then `$PREFIX` would be set to **wor** for **hello**'s\n            autocompletion.\n\nThe expected STDOUT should be an array (list) of any data type. For example:\n\n    [\n        \"Monday\",\n        \"Tuesday\",\n        \"Wednesday\",\n        \"Thursday\",\n        \"Friday\"\n    ]\n    \nYou can additionally include suggestions if any of the array items exactly\nmatches any of the following strings:\n\n* `@IncFiles`   ([read more]((#incfiles-boolean-false)))\n* `@IncDirs`    ([read more]((#incdirs-boolean-false)))\n* `@IncExePath` ([read more]((#incexepath-boolean-false)))\n* `@IncExeAll`  ([read more]((#incexeall-boolean-false)))\n* `@IncManPage` ([read more]((#incmanpage-boolean-false)))\n\n### \"DynamicDesc\": string (zls)\n\nThis is very similar to **Dynamic** except your function should return a\nmap instead of an array. Where each key is the suggestion and the value is\na description.\n\nThe description will appear either in the hint text or alongside the\nsuggestion - depending on which suggestion \"popup\" you define (see\n**ListView**).\n\nTwo variables are created for each **Dynamic** function:\n\n* `ISMETHOD`: `true` if the command being autocompleted is going to run as a\n              pipelined method. `false` if it isn't.\n\n* `PREFIX`: contains the partial term. For example if you typed `hello wor[tab]`\n            then `$PREFIX` would be set to **wor** for **hello**'s\n            autocompletion.\n\nThe expected STDOUT should be an object (map) of any data type. The key is the\nautocompletion suggestion, with the value being the description. For example:\n\n    {\n        \"Monday\": \"First day of the week\",\n        \"Tuesday\": \"Second day of the week\",\n        \"Wednesday\": \"Third day of the week\"\n        \"Thursday\": \"Forth day of the week\",\n        \"Friday\": \"Fifth day of the week\",\n    }\n    \n### \"ExecCmdline\": boolean (false)\n\nSometimes you'd want your autocomplete suggestions to aware of the output\nreturned from the commands that preceded it. For example the suggestions\nfor `[` (index) will depend entirely on what data is piped into it.\n\n**ExecCmdline** tells Murex to run the commandline up until the command\nwhich your cursor is editing and pipe that output to the STDIN of that\ncommands **Dynamic** or **DynamicDesc** code block.\n\n> This is a dangerous feature to enable so **ExecCmdline** is only honoured\n> if the commandline is considered \"safe\". **Dynamic** / **DynamicDesc**\n> will still be executed however if the commandline is \"unsafe\" then your\n> dynamic autocompletion blocks will have no STDIN.\n\nBecause this is a dangerous feature, your partial commandline will only\nexecute if the following conditions are met:\n\n* the commandline must be one pipeline (eg `;` tokens are not allowed)\n* the commandline must not have any new line characters\n* there must not be any redirection, including named pipes\n    (eg `cmd <namedpipe>`) and the STDOUT/STDERR switch token (`?`)\n* the commandline doesn't inline any variables (`$strings`, `@arrays`) or\n    functions (`${subshell}`, `$[index]`)\n* lastly all commands are whitelisted in \"safe-commands\"\n    (`config get shell safe-commands`)\n\nIf these criteria are met, the commandline is considered \"safe\"; if any of\nthose conditions fail then the commandline is considered \"unsafe\".\n\nMurex will come with a number of sane commands already included in its\n`safe-commands` whitelist however you can add or remove them using `config`\n\n    » function: foobar { -> match foobar }\n    » config: eval shell safe-commands { -> append foobar }\n    \nRemember that **ExecCmdline** is designed to be included with either\n**Dynamic** or **DynamicDesc** and those code blocks would need to read\nfrom STDIN:\n\n    autocomplete set \"[\" { [{\n        \"AnyValue\": true,\n        \"AllowMultiple\": true,\n        \"ExecCmdline\": true,\n        \"Dynamic\": ({\n            switch ${ get-type: stdin } {\n                case * {\n                    <stdin> -> [ 0: ] -> format json -> [ 0 ]\n                }\n                \n                catch {\n                    <stdin> -> formap k v { out $k } -> cast str -> append \"]\"\n                }\n            }\n        })\n    }] }\n    \n### \"FileRegexp\": string (zls)\n\nWhen set in conjunction with **IncFiles**, this directive will filter on files\nfiles which match the regexp string. eg to only show \".txt\" extensions you can\nuse the following:\n\n    autocomplete set notepad.exe { [{\n        \"IncFiles\": true,\n        \"FileRegexp\": (\\.txt)\n    }] }\n    \n> Please note that you may need to double escape any regexp strings: escaping\n> the `.` match and then also escaping the escape character in JSON. It is\n> recommended you use the `mxjson` method of quoting using parentheses as this\n> will compile that string into JSON, automatically adding additional escaping\n> where required.\n\n### \"FlagValues\": map of arrays (null)\n\nThis is a map of the flags with the values being the same array of directive\nas the top level.\n\nThis allows you to nest operations by flags. eg when a flag might accept\nmultiple parameters.\n\n**FlagValues** takes a map of arrays, eg\n\n    autocomplete set example { [{\n        \"Flags\": [ \"add\", \"delete\" ],\n        \"FlagValues\": {\n            \"add\": [{\n                \"Flags\": [ \"foo\" ]\n            }],\n            \"delete\": [{\n                \"Flags\": [ \"bar\" ]\n            }]\n        }\n    }] }\n    \n...will provide \"foo\" as a suggestion to `example add`, and \"bar\" as a\nsuggestion to `example delete`.\n\n#### Defaults for matched flags\n\nYou can set default properties to all matched flags by using `*` as a\n**FlagValues** value. To expand the above example...\n\n    autocomplete set example { [{\n        \"Flags\": [ \"add\", \"delete\" ],\n        \"FlagValues\": {\n            \"add\": [{\n                \"Flags\": [ \"foo\" ]\n            }],\n            \"delete\": [{\n                \"Flags\": [ \"bar\" ]\n            }],\n            \"*\": [{\n                \"IncFiles\"\n            }]\n        }\n    }] }\n    \n...in this code we are saying not only does \"add\" support \"foo\" and \"delete\"\nsupports \"bar\", but both \"add\" and \"delete\" also supports any filesystem files.\n\nThis default only applies if there is a matched **Flags** or **FlagValues**.\n\n#### Defaults for any flags (including unmatched)\n\nIf you wanted a default which applied to all **FlagValues**, even when the flag\nwasn't matched, then you can use a zero length string (\"\"). For example\n\n    autocomplete set example { [{\n        \"Flags\": [ \"add\", \"delete\" ],\n        \"FlagValues\": {\n            \"add\": [{\n                \"Flags\": [ \"foo\" ]\n            }],\n            \"delete\": [{\n                \"Flags\": [ \"bar\" ]\n            }],\n            \"\": [{\n                \"IncFiles\"\n            }]\n        }\n    }] }\n    \n### \"Flags\": array of strings (auto-populated from man pages)\n\nSetting **Flags** is the fastest and easiest way to populate suggestions\nbecause it is just an array of strings. eg\n\n    autocomplete set example { [{\n        \"Flags\": [ \"foo\", \"bar\" ]\n    }] }\n    \nIf a command doesn't **Flags** already defined when you request a completion\nsuggestion but that command does have a man page, then **Flags** will be\nautomatically populated with any flags identified from an a quick parse of\nthe man page. However because man pages are written to be human readable\nrather than machine parsable, there may not be a 100% success rate with the\nautomatic man page parsing.\n    \n### \"FlagsDesc\": map of strings (null)\n\nThis is the same concept as **Flags** except it is a map with the suggestion\nas a key and description as a value. This distinction is the same as the\ndifference between **Dynamic** and **DynamicDesc**.\n\nPlease note that currently man page parsing cannot provide a description so\nonly **Flags** get auto-populated.\n\n### \"Goto\": string (zls)\n\nThis is a `goto` in programming terms. While \"ugly\" it does allow for quick and\neasy structural definitions without resorting to writing the entire\nautocomplete in code.\n\n**Goto** takes a string which represents the path to jump to from the top level\nof that autocomplete definition. The path should look something like:\n`/int/string/int/string....` where\n\n* the first character is the separator,\n\n* the first value is an integer that relates to the index in your autocomplete\n    array,\n\n* the second value is a string which points to the flag value map (if you\n    defined **FlagValues**),\n\n* the third value is the integer of the autocomplete array inside that\n**FlagValues** map,\n\n* ...and so on as necessary.\n\nAn example of a really simple **Goto**:\n\n    autocomplete set dd { [\n        {\n            \"Flags\": [ \"if=\", \"of=\", \"bs=\", \"iflag=\", \"oflag=\", \"count=\", \"status=\" ],\n            \"FlagValues\": {\n                \"if\": [{ \n                    \"IncFiles\": true\n                }],\n                \"of\": [{ \n                    \"IncFiles\": true\n                }],\n                \"*\": [{\n                    \"AllowAny\": true\n                }]\n            }\n        },\n        {\n            \"Goto\": \"/0\"\n        }\n    ] }\n    \n**Goto** is given precedence over any other directive. So ensure it's the only\ndirective in it's group.\n\n### \"IgnorePrefix\": boolean (false)\n\nWhen set to `true`, this allows **Dynamic** and **DynamicDesc** functions to\nreturn every result and not just those that match the partial term (as would\nnormally be the default).\n\n### \"IncDirs\": boolean (false)\n\nEnable to include directories.\n\nNot needed if **IncFiles** is set to `true`.\n\nBehavior of this directive can be altered with `config set shell\nrecursive-enabled`\n\n### \"IncExeAll\": boolean (false)\n\nEnable this to any executables. Suggestions will include aliases, functions\nbuiltins and any executables in `$PATH`. It will not include private functions.\n\n### \"IncExePath\": boolean (false)\n\nEnable this to include any executables in `$PATH`. Suggestions will not include\naliases, functions nor privates.\n\n### \"IncFiles\": boolean (true)\n\nInclude files and directories. This is enabled by default for any commands\nthat don't have autocomplete defined but you will need to manually enable\nit in any `autocomplete` schemas you create and want files as part of the\nsuggestions.\n\nIf you want to filter files based on file name then you can set a regexp\nstring to match to using **FileRegexp**.\n\n### \"IncManPage\": boolean (false)\n\nThe default behavior for commands with no autocomplete defined is to parse the\nman page and use those results. If a custom autocomplete is defined then that\nman page parser is disabled by default. You can re-enable it and include its\nresults with other flags and behaviors you define by using this directive.\n\n### \"ListView\": boolean (false)\n\nThis alters the appearance of the autocompletion suggestions \"popup\". Rather\nthan suggestions being in a grid layout (with descriptions overwriting the\nhint text) the suggestions are in a list view with the descriptions next to\nthem on the same row (similar to how an IDE might display it's suggestions).\n\n### \"NestedCommand\": boolean (false)\n\nOnly enable this if the command you are autocompleting is a nested parameter\nof the parent command you have types. For example with `sudo`, once you've\ntyped the command name you wish to elivate, then you would want suggestions\nfor that command rather than for `sudo` itself.\n\n### \"Optional\": boolean (false)\n\nSpecifies if a match is required for the index in this schema. ie optional\nflags.\n\n## See Also\n\n* [`<stdin>` ](../commands/stdin.md):\n  Read the STDIN belonging to the parent code block\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`alias`](../commands/alias.md):\n  Create an alias for a command\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`function`](../commands/function.md):\n  Define a function block\n* [`get-type`](../commands/get-type.md):\n  Returns the data-type of a variable or pipe\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`summary` ](../commands/summary.md):\n  Defines a summary help text for a command\n* [`switch`](../commands/switch.md):\n  Blocks of cascading conditionals\n* [mxjson](../types/mxjson.md):\n  Murex-flavoured JSON (deprecated)"

}
