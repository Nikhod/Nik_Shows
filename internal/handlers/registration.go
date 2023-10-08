package handlers

import (
	"Nik/pkg/helper"
	"Nik/pkg/models"
	"encoding/json"
	"net/http"
)

func (h *Handlers) Registration(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.BadRequest(w, err, h.Logger)
		return
	}

	err = h.Service.IsValidDataForRegistration(&user)
	if err != nil {
		helper.BadRequest(w, err, h.Logger)
		err := helper.ResponseAnswer(w, "Incorrect Name/Login/Password entered for registration!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}
		return
	}

	err = h.Service.RegistrationUser(&user)
	if err != nil {
		err := helper.ResponseAnswer(w, "Login is taken by another user!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}
		return
	}

	err = helper.ResponseAnswer(w, "You have registered successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}
