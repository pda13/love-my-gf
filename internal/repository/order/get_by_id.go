package order

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	"log/slog"
)

func (i *implementation) GetByID(_ context.Context, id uuid.UUID) (domain.Order, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	if order, exists := i.storage[id]; !exists {
		i.logger.Error(errOrderNotFoundByID.Error(), slog.Any("id", id))
		return domain.Order{}, errOrderNotFoundByID
	} else {
		return order, nil
	}
}
