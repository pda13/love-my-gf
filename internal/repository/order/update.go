package order

import (
	"context"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	"log/slog"
)

func (i *implementation) Update(_ context.Context, order domain.Order) error {
	i.mu.Lock()
	i.mu.Unlock()

	if _, exists := i.storage[order.ID]; !exists {
		i.logger.Error(errOrderNotFoundByID.Error(), slog.Any("id", order.ID))
		return errOrderNotFoundByID
	}

	i.storage[order.ID] = order
	return nil
}
