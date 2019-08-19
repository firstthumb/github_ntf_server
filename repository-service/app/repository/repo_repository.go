package repository

import "githubntf/app/entity"

type RepoRepository interface {
	Find(repoID string) (*entity.Repo, error)

	Search(query string) ([]*entity.Repo, error)
}
