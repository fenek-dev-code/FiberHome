package logger

import (
	"log/slog"
	"os"
)

func NewLogger(level string, format string) *slog.Logger {
	// Implementation for creating a new logger based on level and format
	handler := pagesFormatter(format)
	slog.SetDefault(slog.New(handler))
	slog.SetLogLoggerLevel(parseLevel(level))
	return slog.Default()
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func pagesFormatter(format string) slog.Handler {
	// Implementation for custom page formatter
	switch format {
	case "json":
		return slog.NewJSONHandler(os.Stdout, nil)
	case "console":
		return slog.NewTextHandler(os.Stderr, nil)
	default:
		return slog.NewTextHandler(os.Stdout, nil)
	}
}
