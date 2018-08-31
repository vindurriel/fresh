package runner

import (
	"fmt"
	logPkg "log"
	"os"
	"strings"
	"time"
)

type logFunc func(string, ...interface{})

var logger = logPkg.New(os.Stderr, "", 0)

const timeFormat = "15:04:05"

func newLogFunc(prefix string) func(string, ...interface{}) {
	color, clear := "", ""
	if settings["colors"] == "1" {
		color = fmt.Sprintf("\033[%sm", logColor(prefix))
		clear = fmt.Sprintf("\033[%sm", colors["reset"])
	}
	return func(format string, v ...interface{}) {
		now := time.Now()
		timeString := now.Format(timeFormat)
		format = fmt.Sprintf("%s %s%s%s %s", timeString, color, prefix, clear, format)
		logger.Printf(format, v...)
	}
}

func fatal(err error) {
	logger.Fatal(err)
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	for _, line := range strings.Split(string(p), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		appLog(line)
	}
	return len(p), nil
}
