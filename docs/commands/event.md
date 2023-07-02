# `event` - Command Reference

> Event driven programming for shell scripts

## Description

Create or destroy an event interrupt

## Usage

    event: event-type name=interrupt { code block }
    
    !event: event-type name
    
* [Calendar Date Ranges](../mkarray/date.md):
  Create arrays of dates
* [Character arrays](../mkarray/character.md):
  Making character arrays (a to z)
* [Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [Non-Decimal Ranges](../mkarray/non-decimal.md):
  Create arrays of integers from non-decimal number bases
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)

## Examples

Create an event:

    event: onSecondsElapsed autoquit=60 {
        out "You're 60 second timeout has elapsed. Quitting murex"
        exit 1
    }
    
Destroy an event:

    !event onSecondsElapsed autoquit

## Detail

The `interrupt` field in the CLI supports ANSI constants. eg

    event: onKeyPress f1={F1-VT100} {
        tout: qs HintText="Key F1 Pressed"
    }
    
To list compiled event types:

    » runtime --events -> formap event ! { out $event }
    onCommandCompletion
    onFileSystemChange
    onKeyPress
    onPrompt
    onSecondsElapsed

## Synonyms

* `event`
* `!event`


## See Also

* [`formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [`function`](../commands/function.md):
  Define a function block
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`private`](../commands/private.md):
  Define a private function block
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex