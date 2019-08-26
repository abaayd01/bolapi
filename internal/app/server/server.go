package server

import (
	"bolapi/internal/pkg/database"
	"bolapi/internal/pkg/gql"
	"bolapi/internal/pkg/resolvers"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/handler"
	"log"
	"net/http"
)

var defaultPort = "8080"
var port string

func init() {
	err := database.InitDB()
	defer database.CloseDB()

	if err != nil {
		fmt.Println(err)
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
}

func Start() {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
