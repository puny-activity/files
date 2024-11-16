package main

import (
	"context"
	"github.com/puny-activity/files/cfg"
	"github.com/puny-activity/files/internal/app"
	"github.com/puny-activity/files/pkg/lggr"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := cfg.Parse()
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "failed to read config",
			slog.String("error", err.Error()))
		panic(err)
	}

	logger := lggr.New(config.Logger)
	logger.Debug("Logger initialized")

	application := app.New(logger)
	err = application.Start()
	if err != nil {
		logger.LogAttrs(context.Background(), slog.LevelError, "failed to start application",
			slog.String("error", err.Error()))
		panic(err)
	}
	defer func() {
		err := application.Stop()
		if err != nil {
			logger.Error("failed to stop application", "error", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
