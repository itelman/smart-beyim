package handlers

import (
	"log/slog"
	"smart-beyim/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	router.Post("/chat", h.chatPost)
	router.Get("/chat", h.chatGET)

	router.Get("/", h.home)
	return router
}
