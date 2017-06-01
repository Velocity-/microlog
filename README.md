# microlog
Simple logging utility for colourful logging in Go.

## Installing
In your terminal, run `go get github.com/Velocity-/microlog`

## Example
```go
package main

import (
    "github.com/Velocity-/microlog"
)

func main() {
    Trace("Don't mind me, I'm just a %q.", "logger")
    Debug("Starting up...")
    Info("Hey there!")
    Warn("That is about to go down.")
    Error("Oh no, that's not good.")
}
```