package microlog

import (
	"fmt"

	"time"

	"runtime"

	"path"

	"os"

	"github.com/fatih/color"
)

type LogConfig struct {
	IncludeCaller bool
}

var fmtTrace = func(a ...interface{}) string { return fmt.Sprint(a...) }
var fmtDebug = color.New(color.FgGreen).SprintFunc()
var fmtInfo = color.New(color.FgBlue).SprintFunc()
var fmtWarn = color.New(color.FgYellow).SprintFunc()
var fmtError = color.New(color.FgRed).SprintFunc()

var Config LogConfig = LogConfig{
	IncludeCaller: true,
}

func Trace(format string, params ...interface{}) {
	print(format, fmtTrace, "TRACE", params...)
}

func Debug(format string, params ...interface{}) {
	print(format, fmtDebug, "DEBUG", params...)
}

func Info(format string, params ...interface{}) {
	print(format, fmtInfo, "INFO ", params...)
}

func Warn(format string, params ...interface{}) {
	print(format, fmtWarn, "WARN ", params...)
}

func Error(format string, params ...interface{}) {
	print(format, fmtError, "ERROR", params...)
}

func Fatal(format string, params ...interface{}) {
	Error(format, params...)
	os.Exit(0)
}

func Panic(format string, params ...interface{}) {
	Error(format, params...)
	panic(fmt.Sprintf(format, params...))
}

func print(format string, color func(a ...interface{}) string, tag string, params ...interface{}) {
	caller := ""
	if Config.IncludeCaller {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			caller = fmt.Sprintf("%s:%d - ", path.Base(file), line)
		} else {
			caller = "unknown.go:1 - "
		}
	}

	prefix := time.Now().Format("[2006-01-02T15:04:05.000 ") + color(tag) + "] " + caller
	fmt.Printf(prefix+format+"\n", params...)
}
