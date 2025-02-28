package order

import (
	"context"
	business "github.com/pda13/love-my-gf/internal/service/model"
	"github.com/pda13/love-my-gf/pkg/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) ListOrders(ctx context.Context, _ *client.ListOrdersRequest) (*client.ListOrdersResponse, error) {
	resp, err := s.service.ListAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &client.ListOrdersResponse{
		Orders: mapOrders(resp),
	}, nil
}

func mapOrders(in business.Orders) []*client.Order {
	out := make([]*client.Order, 0, len(in))
	for _, o := range in {
		out = append(out, &client.Order{
			Id:       o.ID.String(),
			Item:     o.Item,
			Quantity: o.Quantity,
		})
	}

	return out
}
