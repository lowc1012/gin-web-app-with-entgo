package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lowc1012/gin-web-app-with-entgo/internal/api"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"github.com/urfave/cli/v2"
)

var rootCmd = &cli.App{
	Name:   "myApp",
	Usage:  "Start my golang app",
	Action: startApp,
}

func startApp(*cli.Context) error {
	log.Info("Start MyApp...")
	apiServer, err := api.StartAsync()
	if err != nil {
		log.Fatalw("Failed to start MyApp",
			"error", err.Error(),
		)
		return cli.Exit("Failed to start MyApp", 1)
	}

	// [grateful shutdown]
	// Wait for interrupt signal to shut down the server gracefully with a timeout
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	// blocks app here
	<-quitSignal

	// app waits for 5 secs and closes
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Fatalw("API server shutdown timeout",
			"error", err.Error(),
		)
		return cli.Exit("API server shutdown timeout", 1)
	}

	log.Infow("MyApp shutdown gratefully", "event", "shutdown")
	return nil
}
