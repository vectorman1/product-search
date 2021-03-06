package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vectorman1/product-search/graphql-api/graph"
	"github.com/vectorman1/product-search/graphql-api/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "10002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			}))

	http.Handle("/",
		playground.Handler("GraphQL playground", "/query"),
	)
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
