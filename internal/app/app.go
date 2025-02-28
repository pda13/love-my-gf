package app

import (
	"github.com/pda13/love-my-gf/internal/config"
	orderRepo "github.com/pda13/love-my-gf/internal/repository/order"
	"github.com/pda13/love-my-gf/internal/server"
	orderSvc "github.com/pda13/love-my-gf/internal/service/order"
	"log/slog"
)

type App struct {
	Server *server.Server
}

func New(config *config.Config, log *slog.Logger) (*App, error) {
	repo := orderRepo.New(config, log)
	svc := orderSvc.New(config, log, repo)

	appServer := server.New(server.InitParams{
		Cfg:          config,
		Logger:       log,
		OrderService: svc,
	})

	log.Info("app successfully created")
	return &App{
		Server: appServer,
	}, nil
}
