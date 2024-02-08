package grapgql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/grapgql/graph"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func StartGraphQl(db *gorm.DB) {
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = defaultPort
		}
		router := chi.NewRouter()

		router.Use(middlewares.Middleware())
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Db: db}}))

		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		router.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}()
}
