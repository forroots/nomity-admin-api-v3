package logger

import (
	"log/slog"
	"os"
)

func New(envProduction bool) *slog.Logger {
	var handler slog.Handler

	if envProduction {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug})
	}

	return slog.New(handler)
}
