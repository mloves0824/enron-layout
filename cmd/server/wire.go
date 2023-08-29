//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/mloves0824/enron-layout/internal/biz"
	"github.com/mloves0824/enron-layout/internal/conf"
	"github.com/mloves0824/enron-layout/internal/data"
	"github.com/mloves0824/enron-layout/internal/server"
	"github.com/mloves0824/enron-layout/internal/service"

	"github.com/mloves0824/enron"
	"github.com/mloves0824/enron/log"
	"github.com/google/wire"
)

// wireApp init enron application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*enron.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
