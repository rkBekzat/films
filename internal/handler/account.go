package handler

import (
	"log"
	"net/http"

	"github.com/rkBekzat/films/internal/model"
)

func (h *Handler) registerAccountRoute(router *http.ServeMux) {
	router.HandleFunc("/api/sign_in", h.SignIn)
	router.HandleFunc("/api/sign_up", h.SignUp)
}

// @Summary		Information
// @Security		ApiKeyAuth
// @Tags			Account
// @Accept			json
// @Produce		json
// @Description	Get information current User
// @Router			/api/Account/Me [get]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input signUpInput

	log.Println("enter the sign up: ")
	if err := UnmarshalBody(r.Body, &input); err != nil {
		log.Println("Error: ", err.Error())
		sendErr(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Calling the service: ", input)
	err := h.service.AuthService.CreateUser(&model.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     input.Role,
		Gender:   input.Gender,
		Password: input.Password,
	})
	if err != nil {
		log.Println("Error: ", err.Error())
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	log.Println("Sending response")
	err = sendResponse("ok", w)
	if err != nil {
		log.Println("Error: ", err.Error())
		sendErr(w, http.StatusBadRequest, err)
	}
	log.Println("ok")
}

// @Summary		SignIn
// @Tags			Account
// @Accept			json
// @Produce		json
// @Description	login
// @Param			input	body	signInInput	true	"credentials"
// @Router			/api/Account/SignIn [post]
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	if err := UnmarshalBody(r.Body, &input); err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	log.Println("Input: ", input)
	token, err := h.service.AuthService.GenerateToken(input.Email, input.Password)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
	err = sendResponse(token, w)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}
}
