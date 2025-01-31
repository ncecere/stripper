package crawler

import (
	"fmt"
	"os"
	"time"
)

var debugEnabled = false

func init() {
	debugEnabled = os.Getenv("STRIPPER_DEBUG") == "1"
}

func debugf(format string, args ...interface{}) {
	if debugEnabled {
		timestamp := time.Now().Format("15:04:05.000")
		fmt.Fprintf(os.Stderr, "[DEBUG] %s: %s\n", timestamp, fmt.Sprintf(format, args...))
	}
}
