package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger() {
	env := os.Getenv("ENV")
	var handler slog.Handler

	switch env {
	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	default: // dev
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	Log = slog.New(handler).With("env", env)
}
