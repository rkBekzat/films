package handler

import (
	"net/http"

	"github.com/rkBekzat/films/internal/model"
)

func (h *Handler) registerFilmRoute(router *http.ServeMux) {
	router.HandleFunc("/api/film/add", authorizeMiddlWare(adminAccess(h.AddFilm), h.service.AuthService))
	router.HandleFunc("/api/film/update", authorizeMiddlWare(adminAccess(h.Update), h.service.AuthService))
	router.HandleFunc("/api/film/delete", authorizeMiddlWare(adminAccess(h.Delete), h.service.AuthService))

	router.HandleFunc("/api/film/info", authorizeMiddlWare(h.GetById, h.service.AuthService))
	router.HandleFunc("/api/film/list", authorizeMiddlWare(h.GetFilms, h.service.AuthService))
	router.HandleFunc("/api/film/search", authorizeMiddlWare(h.SearchFilm, h.service.AuthService))
}

func (h *Handler) AddFilm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input model.Film

	if err := UnmarshalBody(r.Body, &input); err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.FilmService.Create(&input)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(id, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	id := params.Get("id")

	res, err := h.service.FilmService.GetById(id)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(res, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}

func (h *Handler) GetFilms(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	offset, limit := params.Get("offset"), params.Get("limit")
	orderBy, order := params.Get("order_by"), params.Get("order")

	res, err := h.service.FilmService.GetFilms(offset, limit, orderBy, order)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(res, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}

func (h *Handler) SearchFilm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	text := params.Get("text")

	res, err := h.service.FilmService.Search(text)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(res, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input model.Film
	params := r.URL.Query()

	if err := UnmarshalBody(r.Body, &input); err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	id := params.Get("id")
	input.Id = id
	err := h.service.FilmService.Update(&input)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(id, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()

	id := params.Get("id")
	err := h.service.FilmService.Delete(id)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(id, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}
