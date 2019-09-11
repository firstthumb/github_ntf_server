package proto

import (
	"context"
	"github.com/twitchtv/twirp"
	pb "githubntf/repository-service-proto/rpc/repository"
	"githubntf/repository-service/app/service"
	"log"
)

type Server struct {
	repoService *service.RepoService
}

func NewProtoServer(s *service.RepoService) *Server {
	return &Server{
		repoService: s,
	}
}

func (s *Server) SearchRepository(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	if req.Query == "" {
		return nil, twirp.RequiredArgumentError("query")
	}

	searchResult, err := s.repoService.Search(req.Query)
	if err != nil {
		log.Printf("Could not search repository : %v", err)
		return nil, twirp.InternalErrorWith(err)
	}

	results := make([]*pb.Result, len(searchResult))
	for index, result := range searchResult {
		results[index] = &pb.Result{
			Name: result.Name,
		}
	}

	return &pb.SearchResponse{
		Results: results,
	}, nil
}
