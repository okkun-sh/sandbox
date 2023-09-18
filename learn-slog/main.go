package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger.InfoContext(context.Background(), "message")

	logger.Info("hoge", slog.Group("req", "method", "POST"))
}
