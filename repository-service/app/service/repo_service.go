package service

import (
	"githubntf/app/entity"
	"githubntf/app/repository"
	"githubntf/app/repository/repo"
	"log"
)

type RepoService struct {
	repoRepository repository.RepoRepository
	db             *repo.MongoRepository
}

func NewRepoService(r repository.RepoRepository, d *repo.MongoRepository) *RepoService {
	return &RepoService{
		repoRepository: r,
		db:             d,
	}
}

func (s *RepoService) Find(repoID string) (*entity.Repo, error) {
	return s.repoRepository.Find(repoID)
}

func (s *RepoService) Search(query string) ([]*entity.Repo, error) {
	cachedSearchResult, err := s.db.FindSearchResult(query)
	if err != nil {
		result, err := s.repoRepository.Search(query)
		if err != nil {
			return nil, err
		}

		err = s.db.StoreSearchResult(query, result)
		if err != nil {
			log.Printf("Could not save search result : %v", err)
		} else {
			log.Printf("[QUERY: %s] => Saving search result to cache", query)
		}

		return result, nil
	}

	log.Printf("[QUERY: %s] => Returned cached search result", query)
	return cachedSearchResult, nil
}
