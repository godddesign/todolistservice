package base

import (
	"io/ioutil"
	"log"
	"os"
)

type (
	Logger interface {
		SetLevel(level string)
		Debug(msg string)
		Debugf(format string, v ...interface{})
		Info(msg string)
		Infof(format string, v ...interface{})
		Error(msg string)
		Errorf(format string, v ...interface{})
		Fatal(err error)
	}

	BaseLogger struct {
		level string
		flag  int
		debug *log.Logger
		info  *log.Logger
		error *log.Logger
	}

	logLevel struct {
		Debug string
		Info  string
		Error string
	}
)

var (
	LogLevel = logLevel{
		Debug: "debug",
		Info:  "info",
		Error: "error",
	}
)

// NewLogger build and return a new logger.
func NewLogger(level string, isFlag bool) *BaseLogger {
	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lmicroseconds //| log.Lshortfile
	}

	return newLogger(level, flag)
}

// SetLevel sets the logging level preference
func newLogger(level string, flag int) *BaseLogger {
	switch level {
	case LogLevel.Debug:
		return &BaseLogger{
			level: LogLevel.Debug,
			flag:  flag,
			debug: log.New(os.Stderr, "DEBUG: ", flag),
			info:  log.New(os.Stderr, "INFO: ", flag),
			error: log.New(os.Stderr, "ERROR: ", flag),
		}

	case LogLevel.Info:
		return &BaseLogger{
			level: LogLevel.Info,
			flag:  flag,
			debug: log.New(ioutil.Discard, "DEBUG: ", flag),
			info:  log.New(os.Stderr, "INFO: ", flag),
			error: log.New(os.Stderr, "ERROR: ", flag),
		}

	case LogLevel.Error:
		return &BaseLogger{
			level: LogLevel.Info,
			flag:  flag,
			debug: log.New(ioutil.Discard, "DEBUG: ", flag),
			info:  log.New(ioutil.Discard, "INFO: ", flag),
			error: log.New(os.Stderr, "ERROR: ", flag),
		}

	default:
		return &BaseLogger{
			level: LogLevel.Info,
			flag:  flag,
			debug: log.New(ioutil.Discard, "DEBUG: ", flag),
			info:  log.New(ioutil.Discard, "INFO: ", flag),
			error: log.New(ioutil.Discard, "ERROR: ", flag),
		}
	}
}

func (bl *BaseLogger) SetLevel(level string) {
	if bl.level != level {
		*bl = *newLogger(level, bl.flag)
	}
}

// Debug - calls l.Output to print to the logger.
func (bl *BaseLogger) Debug(msg string) {
	bl.debug.Println(msg)
}

// Debugf - calls l.Output to print to the logger.
func (bl *BaseLogger) Debugf(format string, v ...interface{}) {
	bl.debug.Printf(format, v...)
}

// Info - calls l.Output to print to the logger.
func (bl *BaseLogger) Info(msg string) {
	bl.info.Println(msg)
}

// Infof - calls l.Output to print to the logger.
func (bl *BaseLogger) Infof(format string, v ...interface{}) {
	bl.info.Printf(format, v...)
}

// Error - calls l.Output to print to the logger.
func (bl *BaseLogger) Error(msg string) {
	bl.error.Println(msg)
}

// Errorf - calls l.Output to print to the logger.
func (bl *BaseLogger) Errorf(format string, v ...interface{}) {
	bl.error.Printf(format, v...)
}

// Dump - calls l.Output to print error to the logger.
func (bl *BaseLogger) Dump(error error) {
	bl.error.Println(error.Error())
}

// Fatal - calls l.Output to print error to the logger and call os.Exit(1).
func (bl *BaseLogger) Fatal(error error) {
	bl.error.Fatal(error.Error())
}
