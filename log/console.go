package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// consoleLogger standard i/o output console logger implement
type consoleLogger struct{}

func (consoleLogger) Debug(v ...interface{}) {
	output(LDebug, fmt.Sprint(v...))
}

func (consoleLogger) Info(v ...interface{}) {
	output(LInfo, fmt.Sprint(v...))
}

func (consoleLogger) Warn(v ...interface{}) {
	output(LWarn, fmt.Sprint(v...))
}

func (consoleLogger) Error(v ...interface{}) {
	output(LError, fmt.Sprint(v...))
}

func (consoleLogger) Debugf(format string, v ...interface{}) {
	output(LDebug, fmt.Sprintf(format, v...))
}

func (consoleLogger) Infof(format string, v ...interface{}) {
	output(LInfo, fmt.Sprintf(format, v...))
}

func (consoleLogger) Warnf(format string, v ...interface{}) {
	output(LWarn, fmt.Sprintf(format, v...))
}

func (consoleLogger) Errorf(format string, v ...interface{}) {
	output(LError, fmt.Sprintf(format, v...))
}

// Flush calls to flush buffer - console logger doesn't need to flush buffer
func (consoleLogger) Flush() error {
	return nil
}

func output(level Level, v ...interface{}) {
	pc, file, line, _ := runtime.Caller(3)
	file = filepath.Base(file)
	funcName := strings.TrimPrefix(filepath.Ext(runtime.FuncForPC(pc).Name()), ".")
	format := "[%s] [%s] %s:%d %s | " + fmt.Sprint(v...) + "\n"
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf(format, level, date, file, line, funcName)
}
