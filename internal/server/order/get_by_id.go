package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/pkg/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetOrder(ctx context.Context, req *client.GetOrderRequest) (*client.GetOrderResponse, error) {
	resp, err := s.service.GetByID(ctx, uuid.MustParse(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &client.GetOrderResponse{
		Order: &client.Order{
			Id:       resp.ID.String(),
			Item:     resp.Item,
			Quantity: resp.Quantity,
		},
	}, nil
}
