package log

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var Level = &slog.LevelVar{}

type Format uint8

const (
	FormatText Format = iota
	FormatJSON
)

func Setup(level slog.Level, format Format, noColor bool) {
	Level.Set(level.Level())
	var h slog.Handler
	if format == FormatJSON {
		h = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: Level}) // nolint: exhaustruct_v5
	} else {
		h = newHandler(os.Stderr, Level.Level(), noColor)
	}
	slog.SetDefault(slog.New(h))
}

func ParseFormat(s string) (Format, error) {
	switch strings.ToLower(s) {
	case "text":
		return FormatText, nil
	case "json":
		return FormatJSON, nil
	default:
		return FormatText, fmt.Errorf("%q is not a valid log format", s)
	}
}

func ParseLevel(s string) (slog.Level, error) {
	switch strings.ToLower(s) {
	case "error":
		return slog.LevelError, nil
	case "warn":
		return slog.LevelWarn, nil
	case "info":
		return slog.LevelInfo, nil
	case "debug":
		return slog.LevelDebug, nil
	default:
		return slog.LevelInfo, fmt.Errorf("%q is not a valid log level", s)
	}
}
