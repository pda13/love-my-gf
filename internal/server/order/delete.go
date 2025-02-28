package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/pkg/client"
)

func (s *server) DeleteOrder(ctx context.Context, req *client.DeleteOrderRequest) (*client.DeleteOrderResponse, error) {
	if err := s.service.Delete(ctx, uuid.MustParse(req.GetId())); err != nil {
		return &client.DeleteOrderResponse{
			Success: false,
		}, nil
	}

	return &client.DeleteOrderResponse{
		Success: true,
	}, nil
}
