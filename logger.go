package pikpakapi

import (
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

func init() {
	debug := os.Getenv("PIKPAKAPI_DEBUG")
	opts := slog.HandlerOptions{}
	if debug == "true" || debug == "1" {
		opts.Level = slog.LevelDebug
	}
	if debug == "2" {
		opts.Level = slog.LevelDebug
		opts.AddSource = true
	}
	logger = slog.New(slog.NewTextHandler(os.Stderr, &opts))
}
