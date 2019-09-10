package client

import (
	pb "githubntf/repository-service-proto/rpc/repository"
	"net/http"
)

func NewRepositoryServiceClient() pb.Github  {
	return pb.NewGithubProtobufClient("http://localhost:8080", &http.Client{})
}