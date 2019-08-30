package server

import (
	"bolapi/internal/pkg/database"
	"bolapi/internal/pkg/gql"
	"bolapi/internal/pkg/resolvers"
	"github.com/99designs/gqlgen/handler"
	"log"
	"net/http"
)

var defaultPort = "8080"

type BolAPIServer struct {
	Port *string
	DB   database.DBInterface
}

func (server *BolAPIServer) Start() {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}})))

	var port string
	if server.Port != nil {
		port = *server.Port
	} else {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
