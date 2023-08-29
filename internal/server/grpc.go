package server

import (
	v1 "github.com/mloves0824/enron-layout/api/helloworld/v1"
	"github.com/mloves0824/enron-layout/internal/conf"
	"github.com/mloves0824/enron-layout/internal/service"

	"github.com/mloves0824/enron/log"
	"github.com/mloves0824/enron/middleware/recovery"
	"github.com/mloves0824/enron/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
