package handlers

import (
	"log/slog"
	"smart-beyim/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Handler struct {
	service *service.Service
	log     *slog.Logger
}

func NewHandler(service *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// Define CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, 
	}

	// Use CORS middleware
	router.Use(cors.New(corsOptions).Handler)

	router.Post("/chat", h.chatPost)
	router.Get("/chat", h.chatGET)

	router.Get("/", h.homeGet)
	return router
}
