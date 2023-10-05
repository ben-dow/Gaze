package logging

import (
	"fmt"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/config"
	"time"
)

const (
	LogLevelTrace = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func Trace(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelTrace {
		log("TRACE", fmt.Sprintf(entry, formats...))
	}
}

func Debug(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelDebug {
		log("DEBUG", fmt.Sprintf(entry, formats...))
	}
}

func Info(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelInfo {
		log("INFO", fmt.Sprintf(entry, formats...))
	}
}

func Warn(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelWarn {
		log("WARN", fmt.Sprintf(entry, formats...))
	}
}

func Error(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelError {
		log("ERROR", fmt.Sprintf(entry, formats...))
	}
}

func log(levelStr string, entry string) {
	fmt.Printf("%s [Gaze] [%s] - %s \n", time.Now().Format(time.RFC3339), levelStr, entry)
}
