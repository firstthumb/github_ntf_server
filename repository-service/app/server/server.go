package server

import (
	"github.com/99designs/gqlgen/handler"
	"githubntf/common/middleware"
	pb "githubntf/repository-service-proto/rpc/repository"
	"githubntf/repository-service/app"
	"githubntf/repository-service/app/config/wire"
	"net/http"
)

type Server interface {
	NewMux() *http.ServeMux
}

type HttpServer struct {
}

func (s *HttpServer) NewMux() *http.ServeMux {
	r := wire.NewRepoService()
	p := wire.NewProtoService()

	twirpHandler := pb.NewGithubServer(p, nil)
	mux := http.NewServeMux()

	mux.Handle("/", handler.Playground("GraphQL playground", "/query"))
	mux.Handle("/query", middleware.CORS(handler.GraphQL(
		app.NewExecutableSchema(
			app.Config{
				Resolvers: r,
			},
		),
	)))
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	return mux
}
