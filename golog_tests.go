package GoLog

import (
	"fmt"
	"os"
)

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
