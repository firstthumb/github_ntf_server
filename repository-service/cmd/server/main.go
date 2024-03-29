package main

import (
	"githubntf/repository-service/app/server"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	mux := (&server.HttpServer{}).NewMux()

	port := getPort()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
