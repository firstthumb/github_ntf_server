package server

import (
	"github.com/99designs/gqlgen/handler"
	"githubntf/app"
	"githubntf/app/config/wire"
	pb "githubntf/proto"
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
	mux.Handle("/query", handler.GraphQL(
		app.NewExecutableSchema(
			app.Config{
				Resolvers: r,
			},
		),
	))
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	return mux
}
