// +build wireinject

package wire

import (
	"github.com/google/wire"
	"githubntf/app/config"
	"githubntf/app/proto"
	"githubntf/app/repository/gateway"
	"githubntf/app/repository/repo"
	"githubntf/app/repository/resolver"
	"githubntf/app/service"
)

var InjectorSet = wire.NewSet(
	config.NewConfig,
	gateway.NewRepoGateway,
	repo.NewMongoRepository,
	service.NewRepoService,
	resolver.NewResolver,
	proto.NewProtoServer,
)

func NewRepoService() *resolver.Resolver {
	panic(wire.Build(InjectorSet))
}

func NewProtoService() *proto.Server {
	panic(wire.Build(InjectorSet))
}
