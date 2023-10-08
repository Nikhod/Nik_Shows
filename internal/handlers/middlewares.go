package handlers

import (
	"Nik/pkg/helper"
	"context"
	"errors"
	"log"
	"net/http"
)

const KeyUserId = "userID"

func (h *Handlers) AdminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//код до Хэндлера
		token := r.Header.Get("token")
		userId, err := h.Service.GetUserIdByToken(token)
		if err != nil {
			helper.BadRequest(w, err, h.Logger)
			return
		}

		isAdmin, err := h.Service.IsAdmin(userId)
		if err != nil {
			helper.Forbidden(w, err, h.Logger)
			return
		}

		if isAdmin == false {
			err = errors.New("Access denied!")
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUserId, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
		// код после Хэндлера
		log.Println("Admin Authentication was Successfuly!")
	})
}

func (h *Handlers) ClientAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cod before the Handler
		token := r.Header.Get("token")
		userId, err := h.Service.GetUserIdByToken(token)
		if err != nil {
			helper.BadRequest(w, err, h.Logger)
			return
		}

		userActive, err := h.Service.GetActiveUserById(userId)
		if err != nil {
			helper.Forbidden(w, err, h.Logger)
			return
		}

		if userActive == false {
			err = helper.ResponseAnswer(w, "The User has been Blocked!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}

		ctx := context.WithValue(r.Context(), KeyUserId, userId)
		r = r.WithContext(ctx)

		log.Println("Client Authentication was Successfuly!")
		next.ServeHTTP(w, r)
		//	Cod after the Handler
		log.Println("Handler has done the work!")

	})
}
