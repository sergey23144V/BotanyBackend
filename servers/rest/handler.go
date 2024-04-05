package rest

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router *chi.Mux) *chi.Mux {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Разрешить запросы от всех источников
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	router.Use(corsMiddleware.Handler)
	router.Post("/save", h.services.SaveImg)

	router.Handle("/img/*", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	router.Handle("/analysis/*", http.StripPrefix("/analysis/", http.FileServer(http.Dir("analysis"))))

	return router
}
