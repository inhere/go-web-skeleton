package clog

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

var debug bool

// SetDebug Set debug
func SetDebug(enable bool) {
	debug = enable
}

// ------- internal console log

// Printf print an message
func Printf(format string, args ...interface{}) {
	color.Println("<green>[INFO] </>", fmt.Sprintf(format, args...))
}

// Debugf print an debug message
func Debugf(format string, args ...interface{}) {
	if debug {
		color.Println("<mga>[DEBUG] </>", fmt.Sprintf(format, args...))
	}
}

// Errorf print an error message
func Errorf(format string, args ...interface{}) {
	color.Println("<red>[ERROR] </>", fmt.Sprintf(format, args...))
}

// Fatalf print an error fatal message an exit
func Fatalf(format string, args ...interface{}) {
	color.Println("<danger>[FATAL] </>", fmt.Sprintf(format, args...))
	os.Exit(1)
}
