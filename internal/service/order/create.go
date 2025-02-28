package order

import (
	"context"
	"github.com/pda13/love-my-gf/internal/logging"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	business "github.com/pda13/love-my-gf/internal/service/model"
)

func (i *implementation) Create(ctx context.Context, order business.Order) (business.Order, error) {
	resp, err := i.repo.Create(ctx, domain.Order{
		ID:       order.ID,
		Item:     order.Item,
		Quantity: order.Quantity,
	})
	if err != nil {
		i.logger.Error("unable to create order", logging.Error(err))
		return business.Order{}, err
	}

	i.logger.Info("order created successfully")
	return business.Order{
		ID:       resp.ID,
		Item:     resp.Item,
		Quantity: resp.Quantity,
	}, nil
}
