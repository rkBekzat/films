package handler

import (
	"net/http"

	"github.com/rkBekzat/films/internal/model"
)

func (h *Handler) registerFilmRoute(router *http.ServeMux) {
	router.HandleFunc("/api/film/add", authorizeMiddlWare(adminAccess(h.AddFilm), h.service.AuthService))
	router.HandleFunc("/api/film/update", authorizeMiddlWare(adminAccess(h.Update), h.service.AuthService))
	router.HandleFunc("/api/film/delete", authorizeMiddlWare(adminAccess(h.DeleteFilm), h.service.AuthService))

	router.HandleFunc("/api/film/info", authorizeMiddlWare(h.GetById, h.service.AuthService))
	router.HandleFunc("/api/film/list", authorizeMiddlWare(h.GetFilms, h.service.AuthService))
	router.HandleFunc("/api/film/search", authorizeMiddlWare(h.SearchFilm, h.service.AuthService))
}

// ShowAccount godoc
// @Summary      Add film
// @Security		ApiKeyAuth
// @Description  adding film
// @Tags         film
// @Accept       json
// @Produce      json
// @Param		input	body	model.Film	true	"user data"
// @Success      200  {object}  string
// @Router       /api/film/add [post]
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

// ShowAccount godoc
// @Summary      get
// @Security		ApiKeyAuth
// @Description  get film by id
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        id   query      string  true  "film ID"
// @Success      200  {object}  model.Film
// @Router       /api/film/info [get]
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

// ShowAccount godoc
// @Summary      get films
// @Security		ApiKeyAuth
// @Description  get films proper
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        offset   query     int  true  "offset"
// @Param        limit   query     int  true  "limit"
// @Param        order_by   query     string  true  "order by which column"
// @Param        order   query     string  true  "order ASC or DESC"
// @Success      200  {object}  []model.Film
// @Router       /api/film/list [get]
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

// ShowAccount godoc
// @Summary      search
// @Security		ApiKeyAuth
// @Description  search film by text
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        text   query     string  true  "search title by text"
// @Success      200  {object}  []model.Film
// @Router       /api/film/search [get]
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

// ShowAccount godoc
// @Summary      update
// @Security		ApiKeyAuth
// @Description  update film by text
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        id   query     string  true  "search title by text"
// @Param		input	body	model.Film	true	"user data"
// @Success      200  {object}  string
// @Router       /api/film/update [put]
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

// @Summary      delete
// @Security		ApiKeyAuth
// @Description  deleting film
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        id   query      string  true  "actor ID"
// @Success      200  {object}  string
// @Router       /api/film/delete [delete]
func (h *Handler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
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
