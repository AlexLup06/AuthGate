package logging

import (
	"fmt"
	"os"
	"strings"

	"log/slog"
)

func New(level string) (*slog.Logger, error) {
	lvl, err := parseLevel(level)
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})

	return slog.New(handler), nil
}

func parseLevel(level string) (slog.Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid log level: %s", level)
	}
}
