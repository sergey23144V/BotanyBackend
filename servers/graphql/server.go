package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const defaultPort = "80"

func StartGraphQl(db *gorm.DB, authServerImpl *implementation.AuthServerImpl, newService *service.Service) {
	go func() {
		port := defaultPort

		router := chi.NewRouter()

		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins: []string{"*"}, // Разрешить запросы от всех источников
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		})
		router.Use(corsMiddleware.Handler)
		router.Use(middlewares.AuthInterceptorGraphQL())
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(db, authServerImpl, newService)}))

		router.Handle("/playground", playground.Handler("GraphQL playground", "/api"))
		router.Handle("/api", srv)

		log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
		log.Printf("connect to http://localhost:%s/api for api", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}()
}
