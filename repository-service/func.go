package githubntf

import (
	"githubntf/repository-service/app/server"
	"net/http"
)

var mux = (&server.HttpServer{}).NewMux()

func RepositoryService(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
