package order

import (
	"context"
	"github.com/pda13/love-my-gf/internal/logging"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	business "github.com/pda13/love-my-gf/internal/service/model"
)

func (i *implementation) Update(ctx context.Context, order business.Order) error {
	if err := i.repo.Update(ctx, domain.Order{
		ID:       order.ID,
		Item:     order.Item,
		Quantity: order.Quantity,
	}); err != nil {
		i.logger.Error("unable to update order", logging.Error(err))
		return err
	}

	i.logger.Info("order updated successfully")
	return nil
}
