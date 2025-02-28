package order

import (
	"context"
	"github.com/google/uuid"
	business "github.com/pda13/love-my-gf/internal/service/model"
	"github.com/pda13/love-my-gf/pkg/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) UpdateOrder(ctx context.Context, req *client.UpdateOrderRequest) (*client.UpdateOrderResponse, error) {
	if err := s.service.Update(ctx, business.Order{
		ID:       uuid.MustParse(req.GetId()),
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &client.UpdateOrderResponse{
		Order: &client.Order{
			Id:       req.GetId(),
			Item:     req.GetItem(),
			Quantity: req.GetQuantity(),
		},
	}, nil
}
