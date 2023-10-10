package handlers

import (
	"Nik/pkg/helper"
	"Nik/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handlers) AddImageToContent(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	err = h.Service.ValidateImage(header.Size)
	if err != nil {
		helper.BadRequest(w, err, h.Logger)
		return
	}

	imageName, err := h.Service.SaveImage(file, header.Filename)
	if err != nil {
		return
	}

	err = h.Service.ValidateImageName(imageName)
	if err != nil {
		helper.BadRequest(w, err, h.Logger)
		return
	}

	err = h.Service.AddImageToDB(imageName, strContentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Image Added successfully!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) AddContentWithLinks(w http.ResponseWriter, r *http.Request) {
	var content models.ContentAndLinks
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		helper.BadRequest(w, err, h.Logger)
		return
	}

	err = h.Service.ValidateContentAndLinks(&content.Content, &content.Links)
	if err != nil {
		helper.ResetContentServerError(w, err, h.Logger)
		return
	}
	log.Println("Validate test")

	vars := mux.Vars(r)
	contentType, genre := vars["content_type"], vars["genre"]

	err = h.Service.AddContent(&content.Content, contentType, genre)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "Maybe this content has already existed...")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}
		return
	}

	err = h.Service.AddLinks(&content.Links, content.Content.Name)
	if err != nil {
		chekErr := errors.New("Content updated!")
		if errors.As(err, &chekErr) {
			helper.ResetContentServerError(w, err, h.Logger)
			err = helper.ResponseAnswer(w, "Content updated!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}

		helper.InternalServerError(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "An error has occurred. Our developers will fix it soon!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}
		return
	}

	err = helper.ResponseAnswer(w, "Content Added Successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) DeleteContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	err := h.Service.DeleteContent(strContentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Content Deleted Successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) BanUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	login := vars["loginKey"]

	err := h.Service.BanUser(login)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "The User is Blocked successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) UnblockUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	login := vars["loginKey"]

	err := h.Service.UnblockUser(login)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "User successfully unlocked!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) DeactivateLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	err := h.Service.DeactivateLink(strContentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Links Deactivated successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) AddRecommendation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	err := h.Service.AddRecomendation(strContentID)
	if err != nil {
		chekErr := errors.New("There is no Content by this name!")
		if errors.As(err, &chekErr) {
			helper.NotFoundErr(w, err, h.Logger)
			err = helper.ResponseAnswer(w, "There is no Content by this name!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}

		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Content Added successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) DeleteFromRecommendation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	err := h.Service.DeleteFromRecomendation(strContentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Content deleted successfuly from recomendation!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

func (h *Handlers) SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Notifications
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		helper.BadRequest(w, err, h.Logger)
		return
	}

	err = h.Service.SendMessage(&message)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Message sent successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}
