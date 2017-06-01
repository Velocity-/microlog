package microlog

import (
	"fmt"
	"strconv"

	"time"

	"runtime"

	"path"

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

func print(format string, color func(a ...interface{}) string, tag string, params ...interface{}) {
	caller := ""
	if Config.IncludeCaller {
		_, file, line, _ := runtime.Caller(2)
		caller = path.Base(file) + ":" + strconv.Itoa(line) + " - "
	}

	prefix := time.Now().Format("[2006-01-02T15:04:05.000 ") + color(tag) + "] " + caller
	fmt.Printf(prefix+format+"\n", params...)
}
