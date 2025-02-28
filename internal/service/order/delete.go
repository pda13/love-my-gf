package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/internal/logging"
)

func (i *implementation) Delete(ctx context.Context, id uuid.UUID) error {
	if err := i.repo.Delete(ctx, id); err != nil {
		i.logger.Error("unable to delete order", logging.Error(err))
		return err
	}

	i.logger.Info("order deleted successfully")
	return nil
}
