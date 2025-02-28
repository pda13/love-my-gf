package order

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/internal/config"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	"log/slog"
	"sync"
)

var _ Repository = &implementation{}

var (
	errOrderNotFoundByID = errors.New("order not found by id")
)

type Repository interface {
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Order, error)
	Update(ctx context.Context, order domain.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListAll(ctx context.Context) (domain.Orders, error)
}

type implementation struct {
	cfg    *config.Config
	logger *slog.Logger

	storage map[uuid.UUID]domain.Order
	mu      *sync.RWMutex
}

func New(cfg *config.Config, logger *slog.Logger) Repository {
	return &implementation{
		storage: make(map[uuid.UUID]domain.Order),
		mu:      new(sync.RWMutex),
		cfg:     cfg,
		logger:  logger,
	}
}
