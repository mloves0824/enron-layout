package server

import (
	v1 "github.com/mloves0824/enron-layout/api/helloworld/v1"
	"github.com/mloves0824/enron-layout/internal/conf"
	"github.com/mloves0824/enron-layout/internal/service"

	"github.com/mloves0824/enron-go/log"
	"github.com/mloves0824/enron-go/middleware/recovery"
	"github.com/mloves0824/enron-go/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
