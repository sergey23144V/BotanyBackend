package rest

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
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

	mime.AddExtensionType(".xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Создаем собственный обработчик файлов
	fileHandler := http.StripPrefix("/analysis/", customFileServer(http.Dir("analysis")))

	// Настраиваем маршрутизатор
	http.Handle("/analysis/", fileHandler)

	router.Use(corsMiddleware.Handler)
	router.Post("/save", h.services.SaveImg)

	router.Handle("/img/*", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	router.Handle("/analysis/*", http.StripPrefix("/analysis/", http.FileServer(http.Dir("analysis"))))

	return router
}

func customFileServer(root http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasSuffix(path, ".xlsx") {
			w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		} else {
			ext := filepath.Ext(path)
			mimeType := mime.TypeByExtension(ext)
			if mimeType != "" {
				w.Header().Set("Content-Type", mimeType)
			}
		}
		http.FileServer(root).ServeHTTP(w, r)
	})
}
