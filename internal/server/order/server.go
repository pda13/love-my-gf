package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/pda13/love-my-gf/internal/config"
	business "github.com/pda13/love-my-gf/internal/service/model"
	"github.com/pda13/love-my-gf/pkg/client"
	"google.golang.org/grpc"
)

var _ client.OrderServiceServer = &server{}

type orderService interface {
	Create(ctx context.Context, order business.Order) (business.Order, error)
	GetByID(ctx context.Context, id uuid.UUID) (business.Order, error)
	Update(ctx context.Context, order business.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListAll(ctx context.Context) (business.Orders, error)
}

type server struct {
	client.UnimplementedOrderServiceServer
	cfg     *config.Config
	service orderService
}

func Register(gRPC *grpc.Server, cfg *config.Config, service orderService) {
	client.RegisterOrderServiceServer(gRPC, &server{
		cfg:     cfg,
		service: service,
	})
}
