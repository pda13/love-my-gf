package order

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (i *implementation) Delete(_ context.Context, id uuid.UUID) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, exists := i.storage[id]; !exists {
		i.logger.Error(errOrderNotFoundByID.Error(), slog.Any("id", id))
		return errOrderNotFoundByID
	}

	delete(i.storage, id)
	return nil
}
