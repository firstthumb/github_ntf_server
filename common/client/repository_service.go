package client

import (
	"githubntf/common/config"
	pb "githubntf/repository-service-proto/rpc/repository"
	"net/http"
)

func NewRepositoryServiceClient(c *config.Config) pb.Github {
	return pb.NewGithubProtobufClient(c.Client.RepositoryServiceUrl, &http.Client{})
}
