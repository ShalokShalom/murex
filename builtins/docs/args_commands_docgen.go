package docs

func init() {

	Definition["args"] = "# _murex_ Shell Docs\n\n## Command Reference: `args` \n\n> Command line flag parser for _murex_ shell scripting\n\n## Description\n\nOne of the nuisances of shell scripts is handling flags. More often than not\nyour script will be littered with `$1` still variables and not handle flags\nshifting in placement amongst a group of parameters. `args` aims to fix that by\nproviding a common tool for parsing flags.\n\n`args` takes a name of a variable to assign the result of the parsed parameters\nas well as a JSON structure containing the result. It also returns a non-zero\nexit number if there is an error when parsing.\n\n## Usage\n\n    args var-name { json-block } -> <stdout>\n\n## Examples\n\n    #!/usr/bin/env murex\n    \n    # First we define what parameters to accept:\n    # Pass the `args` function a JSON string (because JSON objects share the same braces as murex block, you can enter JSON\n    # directly as unescaped values as parameters in murex).\n    #\n    # --str: str == string data type\n    # --num: num == numeric data type\n    # --bool: bool == flag used == true, missing == false\n    # -b: --bool == alias of --bool flag\n    args: args {\n        \"AllowAdditional\": true,\n        \"Flags\": {\n            \"--str\": \"str\",\n            \"--num\": \"num\",\n            \"--bool\": \"bool\",\n            \"-b\": \"--bool\"\n        }\n    }\n    catch {\n        # Lets check for errors in the command line parameters. If they exist then\n        # print the error and then exit.\n        err $args[Error]\n        exit 1\n    }\n    \n    out \"The structure of \\$args is: ${$args->pretty}\\n\\n\"\n    \n    \n    # Some example usage:\n    # -------------------\n    \n    !if { $args->[[ /Flags/--bool ]] } {\n        out \"Flag `--bool` was not set.\"\n    }\n    \n    # `<!null>` redirects the STDERR to a named pipe. In this instance it's the 'null' pipe so equivalent to 2>/dev/null\n    # thus we are just suppressing any error messages.\n    try <!null> {\n        $args -> [[ /Flags/--str ]] -> set fStr\n        $args -> [[ /Flags/--num ]] -> set fNum\n    \n        out \"Defined Flags:\"\n        out \"  --str == $fStr\"\n        out \"  --num == $fNum\"\n    }\n    \n    catch {\n        err \"Missing `--str` and/or `--num` flags.\"\n    }\n    \n    $args -> [ Additional ] -> foreach flag {\n        out \"Additional argument (ie not assigned to a flag): `$flag`.\"\n    }\n\n## See Also\n\n* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by _murex_"

}