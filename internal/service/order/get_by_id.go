package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/internal/logging"
	business "github.com/pda13/love-my-gf/internal/service/model"
)

func (i *implementation) GetByID(ctx context.Context, id uuid.UUID) (business.Order, error) {
	resp, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error("unable to get order by id", logging.Error(err))
		return business.Order{}, err
	}

	i.logger.Info("order found by id successfully")
	return business.Order{
		ID:       resp.ID,
		Item:     resp.Item,
		Quantity: resp.Quantity,
	}, nil
}
