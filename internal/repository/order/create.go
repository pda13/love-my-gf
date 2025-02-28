package order

import (
	"context"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
)

func (i *implementation) Create(_ context.Context, order domain.Order) (domain.Order, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.storage[order.ID] = order
	return order, nil
}
