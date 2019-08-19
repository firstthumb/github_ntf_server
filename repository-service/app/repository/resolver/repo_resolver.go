package resolver

import (
	"context"
	"githubntf/app/entity"
)

type repoResolver struct{}

func (r *repoResolver) ID(ctx context.Context, obj *entity.Repo) (string, error) {
	//return strconv.FormatUint(uint64(obj.ID), 10), nil
	return obj.ID.Hex(), nil
}

func (r *repoResolver) CreatedAt(ctx context.Context, obj *entity.Repo) (string, error) {
	return obj.CreatedAt.String(), nil
}
