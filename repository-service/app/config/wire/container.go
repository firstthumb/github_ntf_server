// +build wireinject

package wire

import (
	"github.com/google/wire"
	"githubntf/common/client"
	"githubntf/common/config"
	"githubntf/repository-service/app/proto"
	"githubntf/repository-service/app/repository/gateway"
	"githubntf/repository-service/app/repository/repo"
	"githubntf/repository-service/app/repository/resolver"
	"githubntf/repository-service/app/service"
)

var InjectorSet = wire.NewSet(
	config.NewConfig,
	gateway.NewRepoGateway,
	repo.NewMongoRepository,
	service.NewRepoService,
	resolver.NewResolver,
	proto.NewProtoServer,
	client.NewRepositoryServiceClient,
)

func NewRepoService() *resolver.Resolver {
	panic(wire.Build(InjectorSet))
}

func NewProtoService() *proto.Server {
	panic(wire.Build(InjectorSet))
}
