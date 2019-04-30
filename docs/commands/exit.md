# _murex_ Language Guide

## Command Reference: `exit`

> Exit murex

### Description

Exit's _murex_ with either a exit number of 0 (by default if no parameters
supplied) or a custom value specified by the first parameter.

### Usage

    exit
    exit number

### Examples

    » exit
    
    » exit 42

### See Also

* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1
* [`null`](../commands/devnull.md):
  null function. Similar to /dev/null