package app

import (
	"github.com/pda13/love-my-gf/internal/api"
	"github.com/pda13/love-my-gf/internal/config"
	orderRepo "github.com/pda13/love-my-gf/internal/repository/order"
	orderSvc "github.com/pda13/love-my-gf/internal/service/order"
	"log/slog"
)

type App struct {
	Server *api.Server
}

func New(config *config.Config, logger *slog.Logger) (*App, error) {
	repo := orderRepo.New(config, logger)
	svc := orderSvc.New(config, logger, repo)

	appServer := api.New(api.InitParams{
		Cfg:          config,
		Logger:       logger,
		OrderService: svc,
	})

	logger.Info("app successfully created")
	return &App{
		Server: appServer,
	}, nil
}
