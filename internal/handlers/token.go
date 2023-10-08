package handlers

import (
	"Nik/pkg/helper"
	"Nik/pkg/models"
	"errors"
	"net/http"
	"time"
)

const (
	loginKey = "login"
	passKey  = "password"
)

func (h *Handlers) GetToken(w http.ResponseWriter, r *http.Request) {
	login, password := r.Header.Get(loginKey), r.Header.Get(passKey)

	isFree := h.Service.Repository.IsLoginFree(login)

	if isFree == true {
		helper.NotFoundErr(w, errors.New("User is not found"), h.Logger)
		err := helper.ResponseAnswer(w, "User is not found! Please register to get Token.")
		if err != nil {
			helper.Forbidden(w, err, h.Logger)
			return
		}
		return
	}

	token, err := h.Service.GenerateToken(login, password)
	if err != nil {
		checkErr := errors.New("User is Blocked!")
		if errors.As(err, &checkErr) {
			helper.InternalServerError(w, err, h.Logger)
			err := helper.ResponseAnswer(w, "User is not found!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}
		helper.BadRequest(w, err, h.Logger)
		return
	}

	// Send Answers
	sendToken := models.SendToken{
		Date:   time.Now(),
		Answer: "Authentication was successful!",
		Token:  token,
	}
	err = helper.SendToken(w, &sendToken)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}
