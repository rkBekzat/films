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
