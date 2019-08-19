package resolver

import (
	"context"
	"githubntf/app/entity"
	"githubntf/app/service"
)

type queryResolver struct {
	repoService *service.RepoService
}

func NewQueryResolver(au *service.RepoService) *queryResolver {
	return &queryResolver{
		repoService: au,
	}
}

func (r *queryResolver) Repo(ctx context.Context, id *string) (*entity.Repo, error) {
	res, err := r.repoService.Find(*id)
	return res, err
}

func (r *queryResolver) Repositories(ctx context.Context, query *string) ([]*entity.Repo, error) {
	res, err := r.repoService.Search(*query)
	return res, err
}
