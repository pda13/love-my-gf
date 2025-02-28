package order

import (
	"context"
	"github.com/pda13/love-my-gf/internal/logging"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	business "github.com/pda13/love-my-gf/internal/service/model"
)

func (i *implementation) ListAll(ctx context.Context) (business.Orders, error) {
	resp, err := i.repo.ListAll(ctx)
	if err != nil {
		i.logger.Error("unable to get all orders list", logging.Error(err))
		return nil, err
	}

	i.logger.Info("list of orders found successfully")
	return mapOrders(resp), nil
}

func mapOrders(in domain.Orders) business.Orders {
	out := make([]business.Order, 0, len(in))
	for _, o := range in {
		out = append(out, business.Order{
			ID:       o.ID,
			Item:     o.Item,
			Quantity: o.Quantity,
		})
	}

	return out
}
