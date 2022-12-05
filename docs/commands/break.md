# _murex_ Shell Docs

## Command Reference: `break`

> terminate execution of a block within your processes scope

## Description

`break` will terminate execution of a block (eg `function`, `private`,
`foreach`, `if`, etc).

`break` requires a parameter and that parameter is the name of the caller
block you wish to break out of. If it is a `function` or `private`, then it
will be the name of that function or private. If it is an `if` or `foreach`
loop, then it will be `if` or `foreach` (respectively).

## Usage

    break block-name

## Examples

    function foo {
        a [1..10] -> foreach i {
            out $i
            if { = i==`5` } then {
                out "exit running function"
                break foo
                out "ended"
            }
        }
    }
    
Running the above code would output:

    » foo
    1
    2
    3
    4
    5
    exit running function

## Detail

`break` cannot escape the bounds of its scope (typically the function it is
running inside). For example, in the following code we are calling `break
bar` (which is a different function) inside of the function `foo`:

    function foo {
        a [1..10] -> foreach i {
            out $i
            if { = i==`5` } then {
                out "exit running function"
                break foo
                out "ended"
            }
        }
    }
    
    function bar {
        foo
    }
    
Regardless of whether we run `foo` or `bar`, both of those functions will
raise the following error:

    Error in `break` ( 7,17): no block found named `bar` within the scope of `foo`

## See Also

* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`die`](../commands/die.md):
  Terminate murex with an exit number of 1
* [commands/`exit`](../commands/exit.md):
  Exit murex
* [commands/`foreach`](../commands/foreach.md):
  Iterate through an array
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`private`](../commands/private.md):
  Define a private function block