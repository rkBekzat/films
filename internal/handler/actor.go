package handler

import (
	"log"
	"net/http"

	"github.com/rkBekzat/films/internal/model"
)

func (h *Handler) registerActorRoute(router *http.ServeMux) {
	router.HandleFunc("/api/actor/add", authorizeMiddlWare(adminAccess(h.AddActor), h.service.AuthService))
	router.HandleFunc("/api/actor/update", authorizeMiddlWare(adminAccess(h.UpdateActorInfo), h.service.AuthService))
	router.HandleFunc("/api/actor/delete", authorizeMiddlWare(adminAccess(h.DeleteActor), h.service.AuthService))

	router.HandleFunc("/api/actor/get", authorizeMiddlWare(h.GetActor, h.service.AuthService))
	router.HandleFunc("/api/actor/search", authorizeMiddlWare(h.SearchActor, h.service.AuthService))
	router.HandleFunc("/api/actor/film_list", authorizeMiddlWare(h.FilmList, h.service.AuthService))
}

// ShowAccount godoc
// @Summary      Add actor
// @Security		ApiKeyAuth
// @Description  adding actor
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param		input	body	model.Actor	true	"user data"
// @Success      200  {object}  string
// @Router       /api/actor/add [post]
func (h *Handler) AddActor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input model.Actor

	if err := UnmarshalBody(r.Body, &input); err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	log.Println("input: ", input)
	id, err := h.service.ActorService.Create(&input)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	sendResponse(id, w)
}

// @Summary      update
// @Security		ApiKeyAuth
// @Description  update actor
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param		input	body	model.Actor	true	"user data"
// @Success      200  {object}  string
// @Router       /api/actor/update [put]
func (h *Handler) UpdateActorInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input model.Actor

	params := r.URL.Query()
	id := params.Get("actor_id")

	if err := UnmarshalBody(r.Body, &input); err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	input.Id = id
	err := h.service.ActorService.Update(&input)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	sendResponse("ok", w)
}

// @Summary      Get
// @Security		ApiKeyAuth
// @Description  get actor
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        actor_id   query      string  true  "actor ID"
// @Success      200  {object}  model.Actor
// @Router       /api/actor/get [get]
func (h *Handler) GetActor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	id := params.Get("actor_id")

	res, err := h.service.ActorService.Read(id)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	sendResponse(res, w)
}

// @Summary      delete
// @Security		ApiKeyAuth
// @Description  deleting actor
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        actor_id   query      string  true  "actor ID"
// @Success      200  {object}  string
// @Router       /api/actor/delete [delete]
func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	id := params.Get("actor_id")

	err := h.service.ActorService.Delete(id)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	sendResponse("ok", w)
}

// @Summary      Add actor
// @Security		ApiKeyAuth
// @Description  adding actor
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        text   query      string  true  "actor ID"
// @Success      200  {object}  []model.Actor
// @Router       /api/actor/search [get]
func (h *Handler) SearchActor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	text := params.Get("text")

	res, err := h.service.ActorService.Search(text)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	sendResponse(res, w)
}

// @Summary      films
// @Security		ApiKeyAuth
// @Description  filmed actor
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        actor_id   query      int  true  "actor ID"
// @Success      200  {object}  []model.Film
// @Router       /api/actor/film_list [get]
func (h *Handler) FilmList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	id := params.Get("actor_id")

	res, err := h.service.ActorService.FilmedList(id)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	sendResponse(res, w)
}
