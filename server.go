package main

import (
	"log"
	"megachasma/db"
	directive "megachasma/graph"
	graph "megachasma/graph/resolver"
	"megachasma/graph/services"
	"megachasma/internal"
	"megachasma/middleware/auth"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db.MigrateDatabase()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := db.NewPostgresConnector()
	service := services.New(db.Conn)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{
		Resolvers: &graph.Resolver{
			Srv: service,
		},
		Directives: directive.Directive,
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.AuthMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
