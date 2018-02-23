# _murex_ command reference

## trypipe

> Checks state of each function in a pipeline and exits block on error


`trypipe` checked the state of each function and exits the pipe if any of them
fail. Where `trypipe` differs from a regular `try` block is that `trypipe` will
check every process along the pipeline as well as the terminating function. The
downside to this is that the piped functions can no longer run in parallel.

    trypipe {
        out: "Hello, World!" -> grep: "non-existent string" -> cat
        out: "This process will be ignored"
    }

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via STDERR than it does via STDOUT

You can see which run mode your functions are executing under via the `fid-list`
command.


### See also

* [trypipe](trypipe): Checks state of each function in a pipeline and exits block on error

* evil
* catch
* fid-list