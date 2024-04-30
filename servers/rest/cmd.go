package rest

import (
	"github.com/go-chi/chi"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"log"
	"net/http"
)

func StartRest(service *service.Service) {
	go func() {
		handlers := NewHandler(service)
		r := chi.NewRouter()

		r.Use(middlewares.AuthInterceptorREST())

		r = handlers.InitRoutes(r)

		// Инициализация сервера
		srv := &http.Server{
			Addr:    ":8080",
			Handler: r, // Использование маршрутизатора chi в качестве обработчика запросов
		}

		log.Println("Starting server on http://localhost:8080")

		// Запуск сервера
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
}
