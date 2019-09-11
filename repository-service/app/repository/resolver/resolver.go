package resolver

import (
	gqlgen_clean_architecture "githubntf/repository-service/app"
	"githubntf/repository-service/app/service"
)

type Resolver struct {
	repoService *service.RepoService
}

func NewResolver(au *service.RepoService) *Resolver {
	return &Resolver{
		repoService: au,
	}
}

func (r Resolver) Repo() gqlgen_clean_architecture.RepoResolver {
	return &repoResolver{}
}

func (r Resolver) Query() gqlgen_clean_architecture.QueryResolver {
	return NewQueryResolver(r.repoService)
}
