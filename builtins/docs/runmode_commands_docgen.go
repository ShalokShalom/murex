package docs

func init() {

	Definition["runmode"] = "# `runmode`\n\n> Alter the scheduler's behaviour at higher scoping level\n\n## Description\n\nDue to dynamic nature in which blocks are compiled on demand, traditional `try`\nand `trypipe` blocks cannot affect the runtime behaviour of schedulers already\ninvoked (eg for function blocks and modules which `try` et al would sit inside).\nTo solve this we need an additional command that is executed by the compiler\nprior to the block being executed which can define the runmode of the scheduler.\nThis is the purpose of `runmode`.\n\nThe caveat of being a compiler command rather than a builtin is that `runmode`\nneeds be the first command in a block.\n\n## Usage\n\n```\nrunmode try|trypipe function|module\n```\n\n## Examples\n\n```\nfunction hello {\n    # Short conversation, exit on error\n    \n    runmode try function\n\n    read name \"What is your name? \"\n    out \"Hello $name, pleased to meet you\"\n    \n    read mood \"How are you feeling? \"\n    out \"I'm feeling $mood too\"\n}\n```\n\n## Detail\n\n`runmode`'s parameters are ordered:\n\n### 1st parameter\n\n#### try\n\nChecks only the last command in the pipeline for errors. However still allows\ncommands in a pipeline to run in parallel.\n\n#### trypipe\n\nChecks every command in the pipeline before executing the next. However this\nblocks pipelines from running every command in parallel.\n\n### 2nd parameter\n\n#### function\n\nSets the runmode for all blocks within the function when `runmode` is placed at\nthe start of the function. This includes privates, autocompletes, events, etc.\n\n#### module\n\nSets the runmode for all blocks within that module when placed at the start of\nthe module. This include any functions, privates, autocompletes, events, etc\nthat are inside that module. The do not need a separate `runmode ... function`\nif `runmode ... module` is set.\n\n## See Also\n\n* [Pipeline](../user-guide/pipeline.md):\n  Overview of what a \"pipeline\" is\n* [Schedulers](../user-guide/schedulers.md):\n  Overview of the different schedulers (or 'run modes') in Murex\n* [`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe`\n* [`event`](../commands/event.md):\n  Event driven programming for shell scripts\n* [`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`function`](../commands/function.md):\n  Define a function block\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`read`](../commands/read.md):\n  `read` a line of input from the user and store as a variable\n* [`try`](../commands/try.md):\n  Handles errors inside a block of code\n* [`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error"

}
