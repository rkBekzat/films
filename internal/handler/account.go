package handler

import "net/http"

func (h *Handler) registerAccountRoute(router *http.ServeMux) {
	router.HandleFunc("/api/sign_in", h.SignIn)
	router.HandleFunc("/api/sign_up", h.SignUp)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

}
