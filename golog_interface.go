package GoLog

import (
	"io"
	"os"
)
	
var logger Logger
var defaultLogger Logger = New(os.Stdout)

func New(writer io.Writer) Logger {
	Init(writer)
	return logger
}

func GetLogger() Logger {

	if logger == nil {
		defaultLogger.Warn("GetLogger called without calling Init(). Returning default logger.", nil)
		return defaultLogger
	}
	
	return logger
}

func Init(writer io.Writer) {
	logger = GoLogger{writer}
}
