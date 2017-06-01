package microlog

import (
	"testing"
)

func TestMicrolog(t *testing.T) {
	Trace("Don't mind me, I'm just a %q.", "logger")
	Debug("Starting up...")
	Info("Hey there!")
	Warn("That is about to go down.")
	Error("Oh no, that's not good.")
}

func TestCleanLog(t *testing.T) {
	Config.IncludeCaller = false
	TestMicrolog(t)
}
