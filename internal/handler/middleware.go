package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rkBekzat/films/internal/service"
)

func authorizeMiddlWare(fnc http.HandlerFunc, auth service.Account) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		log.Println("Token: ", token)
		parts := strings.Split(token, "Bearer ")
		if len(parts) != 2 {
			unauthorized(w, fmt.Errorf("not valide token"))
			return
		}

		jwtToken := parts[1]

		id, role, err := auth.ParseToken(jwtToken)
		if err != nil {
			unauthorized(w, fmt.Errorf("not valid token: %s", err.Error()))
			return
		}

		ctx := context.WithValue(context.WithValue(r.Context(), "user_id", id), "role", role)

		fnc(w, r.WithContext(ctx))
	}
}

func adminAccess(fnc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role := ctx.Value("role").(string)
		if role != "admin" {
			log.Println("role of user: ", role)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user hasn't access: "))
			return
		}
		fnc(w, r)
	}
}

func unauthorized(w http.ResponseWriter, err error) {
	log.Println("Error: ", err.Error())
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized: " + err.Error()))
}
