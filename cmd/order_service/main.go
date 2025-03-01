package main

import (
	"github.com/joho/godotenv"
	"github.com/pda13/love-my-gf/internal/app"
	"github.com/pda13/love-my-gf/internal/config"
	"github.com/pda13/love-my-gf/internal/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := logging.New()
	logger.Info("logger successfully initialized")

	if err := godotenv.Load(); err != nil {
		logger.Info("error while loading env variables:", logging.Error(err))
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Error("unable to load config", logging.Error(err))
		os.Exit(1)
	}

	logger.Info("config loaded successfully")

	application, err := app.New(cfg, logger)
	if err != nil {
		logger.Error("error occurred while creating an app", logging.Error(err))
		os.Exit(1)
	}

	go func(app *app.App) {
		if err := app.Server.Run(); err != nil {
			logger.Error("error occurred while running a gRPC api", logging.Error(err))
			os.Exit(1)
		}
	}(application)

	stopSignalChannel := make(chan os.Signal, 1)
	signal.Notify(stopSignalChannel, syscall.SIGTERM, syscall.SIGINT)
	<-stopSignalChannel

	application.Server.GracefulStop()
	logger.Info("api gracefully stopped")
}
