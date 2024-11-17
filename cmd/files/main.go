package main

import (
	"github.com/puny-activity/files/cfg"
	"github.com/puny-activity/files/internal/app"
	"github.com/puny-activity/files/pkg/lggr"
	"github.com/puny-activity/files/pkg/werr"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := cfg.Parse()
	if err != nil {
		panic(werr.Wrap("failed to read config", err))
	}

	logger := lggr.New(config.Logger)
	logger.Debug("Logger initialized")

	application := app.New(logger)
	err = application.Start()
	if err != nil {
		panic(werr.Wrap("failed to start application", err))
	}
	defer func() {
		err := application.Stop()
		if err != nil {
			logger.Warn("failed to stop application", "error", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
