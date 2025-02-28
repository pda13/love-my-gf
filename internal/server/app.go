package server

import (
	"errors"
	"fmt"
	"github.com/pda13/love-my-gf/internal/config"
	grpcmw "github.com/pda13/love-my-gf/internal/middleware/grpc"
	orderServer "github.com/pda13/love-my-gf/internal/server/order"
	"github.com/pda13/love-my-gf/internal/service/order"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

var (
	errStartingTcpConnection = errors.New("unable to start tcp connection")
	errStartingServer        = errors.New("unable to start gRPC server")
)

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
	server *grpc.Server
}

type InitParams struct {
	Cfg          *config.Config
	Logger       *slog.Logger
	OrderService order.Service
}

func New(initParams InitParams) *Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmw.ValidationUnaryServerInterceptor(),
		),
	)

	orderServer.Register(server, initParams.Cfg, initParams.OrderService)
	return &Server{
		cfg:    initParams.Cfg,
		logger: initParams.Logger,
		server: server,
	}
}

func (server *Server) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.cfg.GrpcServer.Port))
	if err != nil {
		return errStartingTcpConnection
	}

	server.logger.Info("gRPC server is running", slog.String("address", listener.Addr().String()))
	if err := server.server.Serve(listener); err != nil {
		return errStartingServer
	}

	return nil
}

func (server *Server) GracefulStop() {
	server.logger.Info("gracefully stopping gRPC server...")
	server.server.GracefulStop()
}
