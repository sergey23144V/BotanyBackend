package grapgql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/auth"
	"github.com/sergey23144V/BotanyBackend/servers/grapgql/graph"
	"log"
	"net/http"
)

const defaultPort = "8080"

func StartGraphQl(db *gorm.DB, authServerImpl *auth.AuthServerImpl, newService *service.Service) {
	go func() {
		port := defaultPort

		router := chi.NewRouter()

		router.Use(middlewares.AuthInterceptorGraphQL())
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(db, authServerImpl, newService)}))

		router.Handle("/playground", playground.Handler("GraphQL playground", "/api"))
		router.Handle("/api", srv)

		log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
		log.Printf("connect to http://localhost:%s/api for api", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}()
}
