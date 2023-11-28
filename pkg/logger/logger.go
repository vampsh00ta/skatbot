package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	runLogFile, _ := os.OpenFile(
		"myapp.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	logger := zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
	return &Logger{
		&logger,
	}
}
