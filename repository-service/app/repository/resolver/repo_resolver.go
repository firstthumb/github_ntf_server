package resolver

import (
	"context"
	"githubntf/repository-service/app/entity"
	"strconv"
)

type repoResolver struct{}

/*
func (r *repoResolver) ID(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.ID.Hex(), nil
}
*/
func (r *repoResolver) RepoID(ctx context.Context, obj *entity.Repo) (string, error) {
	return strconv.FormatInt(obj.RepoId, 10), nil
}

/*
func (r *repoResolver) Name(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.Name, nil
}

func (r *repoResolver) Owner(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.Owner, nil
}

func (r *repoResolver) OwnerId(ctx context.Context, obj *entity.Repo) (int64, error) {
	return obj.OwnerId, nil
}

func (r *repoResolver) Description(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.Description, nil
}

func (r *repoResolver) Language(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.Language, nil
}
*/
func (r *repoResolver) CreatedAt(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.CreatedAt.String(), nil
}
