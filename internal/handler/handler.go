package handler

import (
	"net/http"

	_ "github.com/rkBekzat/films/docs"
	"github.com/rkBekzat/films/internal/service"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	h.registerAccountRoute(router)

	router.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
