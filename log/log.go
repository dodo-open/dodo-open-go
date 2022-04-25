package log

// Logger logger
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})

	// Flush calls to flush buffer
	Flush() error
}

type Level string

const (
	LDebug Level = "Debug"
	LInfo  Level = "Info"
	LWarn  Level = "Warn"
	LError Level = "Error"
)

var DefaultLogger = Logger(new(consoleLogger))

// Debug log (level Debug)
func Debug(v ...interface{}) {
	DefaultLogger.Debug(v...)
}

// Info log (level Info)
func Info(v ...interface{}) {
	DefaultLogger.Info(v...)
}

// Warn log (level Warn)
func Warn(v ...interface{}) {
	DefaultLogger.Warn(v...)
}

// Error log (level Error)
func Error(v ...interface{}) {
	DefaultLogger.Error(v...)
}

// Debugf log (level Debug)
func Debugf(format string, v ...interface{}) {
	DefaultLogger.Debugf(format, v...)
}

// Infof log (level Info)
func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

// Warnf log (level Warn)
func Warnf(format string, v ...interface{}) {
	DefaultLogger.Warnf(format, v...)
}

// Errorf log (level Error)
func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

// Flush flush buffer
func Flush() {
	_ = DefaultLogger.Flush()
}
