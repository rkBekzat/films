package handler

import (
	"net/http"

	"github.com/rkBekzat/films/internal/service"
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
}
