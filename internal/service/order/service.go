package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/internal/config"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	business "github.com/pda13/love-my-gf/internal/service/model"
	"log/slog"
)

var _ Service = &implementation{}

type (
	Service interface {
		Create(ctx context.Context, order business.Order) (business.Order, error)
		GetByID(ctx context.Context, id uuid.UUID) (business.Order, error)
		Update(ctx context.Context, order business.Order) error
		Delete(ctx context.Context, id uuid.UUID) error
		ListAll(ctx context.Context) (business.Orders, error)
	}

	repository interface {
		Create(ctx context.Context, order domain.Order) (domain.Order, error)
		GetByID(ctx context.Context, id uuid.UUID) (domain.Order, error)
		Update(ctx context.Context, order domain.Order) error
		Delete(ctx context.Context, id uuid.UUID) error
		ListAll(ctx context.Context) (domain.Orders, error)
	}
)

type implementation struct {
	cfg    *config.Config
	logger *slog.Logger
	repo   repository
}

func New(cfg *config.Config, logger *slog.Logger, repo repository) Service {
	return &implementation{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}
}
