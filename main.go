package main

import (
	"fmt"
	"os"
	"io"
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

type Logger interface {
	Log(s string, params ...interface{})
	Debug(s string, params ...interface{})
	Warn(s string, err error, params ...interface{})
	Error(s string, err error, params ...interface{})
}

type GoLogger struct {
	writer  io.Writer
}

func (logger GoLogger) Log(s string, params ...interface{}) {
	s = FG_GREEN + s + RESET;
	logger._print(s, params...)
}

func (logger GoLogger) Debug(s string, params ...interface{}) {
	s = FG_CYAN + s + RESET;
	logger._print(s, params...)
}

func (logger GoLogger) Warn(s string, err error, params ...interface{}) {
	s = FG_MAGENTA + s + RESET;
	logger._print(s, params...)
}

func (logger GoLogger) Error(s string, err error, params ...interface{}) {

	if (err == nil) {
		s = fmt.Sprintf(FG_RED + s + RESET, params...)
	} else {
		s = fmt.Sprintln(fmt.Sprintf(FG_RED + s, params...))
		s += fmt.Sprintln("Cause: " + err.Error() + RESET)
	}

	logger._print(s)
}

func (logger GoLogger) _print(s string, params ...interface{}) {
	if (len(params) == 0) {
		fmt.Fprintln(logger.writer, s)
	} else {
		fmt.Fprintln(logger.writer, fmt.Sprintf(s, params...))
	}
}

func New(writer io.Writer) Logger {
	return GoLogger{writer}
}



func main() {
	loggerTest()
	fmt.Println()
	colorTest()
}

func loggerTest() {
	logger := New(os.Stdout)
	logger.Log("Hello %s!", "Log")
	logger.Debug("Hello Debug!")
	logger.Warn("Hello Warn!", nil)
	logger.Error("Hello Error!", nil)
}

func colorTest() {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_RED,     "Hello, Red Sailor!",     RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_GREEN,   "Hello, Green Sailor!",   RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_BLUE,    "Hello, Blue Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_YELLOW,  "Hello, Yellow Sailor!",  RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_CYAN,    "Hello, Cyan Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_MAGENTA, "Hello, Magenta Sailor!", RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_WHITE,   "Hello, White Sailor!",   RESET)
	fmt.Println();
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_RED_BRIGHT,     "Hello, Red Sailor!",     RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_GREEN_BRIGHT,   "Hello, Green Sailor!",   RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_BLUE_BRIGHT,    "Hello, Blue Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_YELLOW_BRIGHT,  "Hello, Yellow Sailor!",  RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_CYAN_BRIGHT,    "Hello, Cyan Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_MAGENTA_BRIGHT, "Hello, Magenta Sailor!", RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", FG_WHITE_BRIGHT,   "Hello, White Sailor!",   RESET)
	fmt.Println();
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_RED,     "Hello, Red Sailor!",     RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_GREEN,   "Hello, Green Sailor!",   RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_BLUE,    "Hello, Blue Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_YELLOW,  "Hello, Yellow Sailor!",  RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_CYAN,    "Hello, Cyan Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_MAGENTA, "Hello, Magenta Sailor!", RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_WHITE,   "Hello, White Sailor!",   RESET)
	fmt.Println();
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_RED_BRIGHT,     "Hello, Red Sailor!",     RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_GREEN_BRIGHT,   "Hello, Green Sailor!",   RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_BLUE_BRIGHT,    "Hello, Blue Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_YELLOW_BRIGHT,  "Hello, Yellow Sailor!",  RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_CYAN_BRIGHT,    "Hello, Cyan Sailor!",    RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_MAGENTA_BRIGHT, "Hello, Magenta Sailor!", RESET)
	fmt.Fprintf(os.Stdout, "%s%s%s\n", BG_WHITE_BRIGHT,   "Hello, White Sailor!",   RESET)
}

