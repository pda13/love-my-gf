package order

import (
	"context"
	"github.com/google/uuid"
	business "github.com/pda13/love-my-gf/internal/service/model"
	"github.com/pda13/love-my-gf/pkg/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateOrder(ctx context.Context, req *client.CreateOrderRequest) (*client.CreateOrderResponse, error) {
	resp, err := s.service.Create(ctx, business.Order{
		ID:       uuid.New(),
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &client.CreateOrderResponse{
		Id: resp.ID.String(),
	}, nil
}
