package logging

import (
	"bytes"
	"fmt"
	"github.com/ben-dow/Gaze/cmd/gaze/svc/config"
	"log"
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
		baseLog("TRACE", fmt.Sprintf(entry, formats...))
	}
}

func Debug(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelDebug {
		baseLog("DEBUG", fmt.Sprintf(entry, formats...))
	}
}

func Info(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelInfo {
		baseLog("INFO", fmt.Sprintf(entry, formats...))
	}
}

func Warn(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelWarn {
		baseLog("WARN", fmt.Sprintf(entry, formats...))
	}
}

func Error(entry string, formats ...any) {
	if config.GetConfiguration().LogLevel <= LogLevelError {
		baseLog("ERROR", fmt.Sprintf(entry, formats...))
	}
}

func baseLog(levelStr string, entry string) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("[Gaze] [%s] ", levelStr))
	space := []byte(" ")
	buffer.Write(bytes.Repeat(space, 17-buffer.Len()))
	buffer.WriteString(fmt.Sprintf("%s", entry))

	log.Println(buffer.String())
}
