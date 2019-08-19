package gateway

import (
	"context"
	"errors"
	"github.com/google/go-github/v27/github"
	"githubntf/app/entity"
	"githubntf/app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type repoGateway struct {
	client *github.Client
}

func NewRepoGateway() repository.RepoRepository {
	return &repoGateway{
		client: github.NewClient(nil),
	}
}

func (r *repoGateway) Search(query string) ([]*entity.Repo, error) {
	log.Printf("Requesting search with query : %s", query)
	result, _, err := r.client.Search.Repositories(context.TODO(), query, &github.SearchOptions{})
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, errors.New("could not search repository")
	}

	log.Printf("Found result : %d ", len(result.Repositories))

	repos := make([]*entity.Repo, len(result.Repositories))
	for index, rp := range result.Repositories {
		repos[index] = &entity.Repo{
			Name: *rp.Name,
		}
	}

	return repos, nil
}

func (r *repoGateway) Find(repoID string) (*entity.Repo, error) {
	return &entity.Repo{
		ID:        primitive.NewObjectID(),
		Name:      "Name",
		CreatedAt: time.Time{},
	}, nil
}
