package handlers

import (
	"Nik/internal/services"
	"Nik/pkg/helper"
	"Nik/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handlers struct {
	Service *services.Service
	Logger  *zap.Logger
}

func NewHandler(service *services.Service, logger *zap.Logger) *Handlers {
	return &Handlers{Service: service, Logger: logger}
}

func (h *Handlers) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	err = h.Service.DeleteMyAccount(userID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Account deleted successfully!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) Support(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	userName, err := h.Service.Repository.GetUserNameByID(userID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	supportText := userName + ", If you have any suggestions or complaints, " +
		"you can write to us by email and we will definitely improve:"

	supportStruct := models.Support{
		SupportText: supportText,
		Email:       "support_nik_show@gmail.com",
	}
	supportByte, err := json.MarshalIndent(supportStruct, "", " ")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	_, err = w.Write(supportByte)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
	return
}

//  Message's handlers Message's handlers Message's handlers Message's handlers Message's handlers

func (h *Handlers) ViewAllMessages(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strPage, strCount := vars["page"], vars["count"]

	messages, err := h.Service.ViewAllMessages(userID, strPage, strCount)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendMessages(w, messages, h.Logger)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) ViewUnreadMessages(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strPage, strCount := vars["page"], vars["count"]

	messages, err := h.Service.ViewUnreadMessages(userID, strPage, strCount)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendMessages(w, messages, h.Logger)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) DeleteMessageById(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strNotificationID := vars["notification_id"]

	err = h.Service.DeleteMessageById(userID, strNotificationID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Message deleted Successfuly!")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) ReadMessage(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strNotificationID := vars["notification_id"]

	message, err := h.Service.ReadMessage(userID, strNotificationID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendMessages(w, message, h.Logger)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) MarkAllMessagesAsRead(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	err = h.Service.MarkAsReadAllMessages(userID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Messages marked as read ✔")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) MarkTheMessageAsRead(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strNotificationID := vars["notification_id"]

	err = h.Service.MarkTheMessageAsRead(userID, strNotificationID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "The Message marked ✔")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) UnmarkTheMessageAsRead(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strNotificationID := vars["notification_id"]

	err = h.Service.UnmarkTheMessageAsRead(userID, strNotificationID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Message Unmarked ✔")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

// playlist's handlers playlist's handlers playlist's handlers playlist's handlers playlist's handlers

func (h *Handlers) ViewUserRecommendations(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strPage, strCount := vars["page"], vars["count"]

	recommendations, err := h.Service.ViewUserRecommendations(userID, strPage, strCount)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContents(w, recommendations)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) ViewAdminRecommendations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strPage, strCount := vars["page"], vars["count"]

	recommendations, err := h.Service.ViewAdminRecommendations(strPage, strCount)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContents(w, recommendations)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) ViewPremieres(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strPage, strCount := vars["page"], vars["count"]

	premieres, err := h.Service.ViewPremieres(strPage, strCount)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContents(w, premieres)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) ViewAllContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strPage, strCount := vars["page"], vars["count"]

	contents, err := h.Service.ViewAllContent(strPage, strCount)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContents(w, contents)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) AddToMyPlaylists(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strContentID, playlist := vars["content_id"], vars["playlist"]

	err = h.Service.AddToMyPlaylist(userID, strContentID, playlist)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		chekErr := errors.New("Сontent type does not match")
		if err != chekErr {
			helper.BadRequest(w, err, h.Logger)
			err = helper.ResponseAnswer(w, "You cannot add this content to 'MyMovies'!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}
		return
	}

	err = helper.ResponseAnswer(w, "Content added successfully")
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}
}

func (h *Handlers) ViewFromMyPlaylist(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strPage, strCount, playlist := vars["page"], vars["count"], vars["playlist"]

	myMovies, err := h.Service.ViewContentsFromMyPlaylist(userID, strPage, strCount, playlist)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContents(w, myMovies)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) DeleteContentFromPlaylist(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strContentID, playlist := vars["content_id"], vars["playlist"]

	err = h.Service.DeleteContentFromPlaylist(userID, strContentID, playlist)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.ResponseAnswer(w, "Content Deleted successfully From "+playlist)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) ViewConcreteContentByID(w http.ResponseWriter, r *http.Request) {
	userID, err := helper.GetUserIdFromContext(r)
	if err != nil {
		helper.NotFoundErr(w, err, h.Logger)
		err = helper.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helper.InternalServerError(w, err, h.Logger)
			return
		}

		return
	}

	vars := mux.Vars(r)
	strContent := vars["content_id"]

	contentID, err := strconv.Atoi(strContent)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	content, err := h.Service.ViewConcreteContentByID(userID, contentID)
	if err != nil {
		checkErr := errors.New("You haven't reached a certain age yet!")
		if errors.As(err, &checkErr) {
			helper.InternalServerError(w, err, h.Logger)
			err = helper.ResponseAnswer(w, "You haven't reached a certain age yet!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}

		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContent(w, content)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) GetLinks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	links, err := h.Service.GetLinks(contentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendLinks(w, links)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) GetContentByFilter(w http.ResponseWriter, r *http.Request) {
	var filter models.Filter
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	content, err := h.Service.GetContentsByFilter(&filter)
	if err != nil {
		checkErr := errors.New("Your request does not meet the conditions!")
		if errors.As(err, &checkErr) {
			helper.BadRequest(w, err, h.Logger)
			err = helper.ResponseAnswer(w, "Your request does not meet the conditions!")
			if err != nil {
				helper.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	err = helper.SendContents(w, content)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}

func (h *Handlers) GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strContentID := vars["content_id"]

	image, err := h.Service.GetImage(strContentID)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_, err = w.Write(image)
	if err != nil {
		helper.InternalServerError(w, err, h.Logger)
		return
	}

	return
}
