package GoLog

import (
	"fmt"
	"os"
	"io"
	"runtime"
	"strings"
	"time"
)

const (
	FG_BLACK   = "\u001b[30m"
	FG_RED     = "\u001b[31m"
	FG_GREEN   = "\u001b[32m"
	FG_YELLOW  = "\u001b[33m"
	FG_BLUE    = "\u001b[34m"
	FG_MAGENTA = "\u001b[35m"
	FG_CYAN    = "\u001b[36m"
	FG_WHITE   = "\u001b[37m")

const (
	FG_BLACK_BRIGHT   = "\u001b[90m"
	FG_RED_BRIGHT	  = "\u001b[91m"
	FG_GREEN_BRIGHT	  = "\u001b[92m"
	FG_YELLOW_BRIGHT  = "\u001b[93m"
	FG_BLUE_BRIGHT	  = "\u001b[94m"
	FG_MAGENTA_BRIGHT = "\u001b[95m"
	FG_CYAN_BRIGHT	  = "\u001b[96m"
	FG_WHITE_BRIGHT   = "\u001b[97m")

const (
	BG_BLACK   = "\u001b[40m"
	BG_RED	   = "\u001b[41m"
	BG_GREEN   = "\u001b[42m"
	BG_YELLOW  = "\u001b[43m"
	BG_BLUE	   = "\u001b[44m"
	BG_MAGENTA = "\u001b[45m"
	BG_CYAN	   = "\u001b[46m"
	BG_WHITE   = "\u001b[47m")

const (
	BG_BLACK_BRIGHT   = "\u001b[100m"
	BG_RED_BRIGHT	  = "\u001b[101m"
	BG_GREEN_BRIGHT	  = "\u001b[102m"
	BG_YELLOW_BRIGHT  = "\u001b[103m"
	BG_BLUE_BRIGHT	  = "\u001b[104m"
	BG_MAGENTA_BRIGHT = "\u001b[105m"
	BG_CYAN_BRIGHT	  = "\u001b[106m"
	BG_WHITE_BRIGHT   = "\u001b[107m")

const (
	RESET      = "\u001b[0m"
	BOLD       = "\u001b[1m"
	FAINT      = "\u001b[2m"
	ITALIC     = "\u001b[3m"
	UNDERLINE  = "\u001b[4m"
	SLOW_BLINK = "\u001b[5m"
	FAST_BLINK = "\u001b[6m"
	SWAP       = "\u001b[7m"
	CONCEAL    = "\u001b[8m"
	CROSS_OUT  = "\u001b[9m"
)

const TIME_LAYOUT = "01/02/06 15:04:05 MST"

var gopath = os.Getenv("GOPATH")

type Logger struct {
	writer  io.Writer
}

func (logger *Logger) Log(s string, params ...interface{}) {
	_log(logger, s, "LOG", FG_GREEN, params...)
}

func (logger *Logger) Debug(s string, params ...interface{}) {
	_log(logger, s, "DEBUG", FG_CYAN, params...)
}

func (logger *Logger) Warn(s string, err error, params ...interface{}) {
	_error(logger, s, "WARN", FG_MAGENTA, err, params...);
}

func (logger *Logger) Error(s string, err error,  params ...interface{}) {
	_error(logger, s, "ERROR", FG_RED, err, params...);
}

func _log(logger *Logger, s string, t string, color string, params ...interface{}) {
	tim, fileName, lineNumber := _getLogData();
	
	fmt.Fprintf(logger.writer,
		"%s%s (%s[%d]) [%s] %s%s\n",
		color,
		tim.Format(TIME_LAYOUT),
		fileName,
		lineNumber,
		t,
		fmt.Sprintf(s, params...),
		RESET);
}

func _error(logger *Logger, s string, t string, color string, err error, params ...interface{}) {
	tim, fileName, lineNumber := _getLogData();
	
	if err == nil {
		fmt.Fprintf(logger.writer,
			"%s%s (%s[%d]) [%s] %s%s\n",
			color,
			tim.Format(TIME_LAYOUT),
			fileName,
			lineNumber,
			t,
			fmt.Sprintf(s, params...),
			RESET);
	} else {
		fmt.Fprintf(logger.writer,
			"%s%s (%s[%d]) [%s] %s\n\tCause: %s%s\n",
			color,
			tim.Format(TIME_LAYOUT),
			fileName,
			lineNumber,
			t,
			fmt.Sprintf(s, params...),
			err.Error(),
			RESET);
	}
}

func _getLogData() (time.Time, string, int) {
	_, fileName, lineNumber, ok := runtime.Caller(3)

	if !ok {
		fileName = "FILENAME NOT RECOVERABLE"
		lineNumber = 0
	} else if strings.Contains(fileName, gopath) {
		fileName = strings.Replace(fileName, gopath, "$GOPATH", 1)
	}
	
	return time.Now(), fileName, lineNumber
}
