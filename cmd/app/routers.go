package main

import (
	"Nik/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitMux(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/registration", h.Registration).Methods(http.MethodPost)
	router.HandleFunc("/sign-in", h.GetToken).Methods(http.MethodGet)

	// Admin's Handlers
	admin := router.PathPrefix("/admin").Subrouter()
	admin.Use(h.AdminAuthentication)

	admin.HandleFunc("/content", h.AddContentWithLinks).Queries(
		"content_type", "{content_type}", "genre", "{genre}").Methods(http.MethodPost)

	admin.HandleFunc("/image/{content_id}", h.AddImageToContent).Methods(http.MethodPost)

	admin.HandleFunc("/recommendation/{content_id}", h.AddRecommendation).Methods(http.MethodPost)

	admin.HandleFunc("/messages", h.SendMessage).Methods(http.MethodPost)

	admin.HandleFunc("/unblock/user", h.UnblockUser).Queries(
		"login", "{login}").Methods(http.MethodPatch)

	// Admin Block -------------------------------------------------------------------------------------------
	block := router.PathPrefix("/admin/delete").Subrouter()
	block.Use(h.AdminAuthentication)

	block.HandleFunc("/user", h.BanUser).Queries(
		"login", "{login}").Methods(http.MethodDelete)

	block.HandleFunc("/recommendation/{content_id}", h.DeleteFromRecommendation).Methods(http.MethodDelete)

	block.HandleFunc("/content/{content_id}", h.DeleteContent).Methods(http.MethodDelete)

	block.HandleFunc("/link/{content_id}", h.DeactivateLink).Methods(http.MethodDelete)

	// Client's Handlers !!! Client's Handlers !!! Client's Handlers !!! Client's Handlers !!! Client's Handlers!!!
	client := router.PathPrefix("/client").Subrouter()
	client.Use(h.ClientAuthentication)

	client.HandleFunc("/account", h.DeleteAccount).Methods(http.MethodDelete)
	client.HandleFunc("/support", h.Support).Methods(http.MethodGet)
	client.HandleFunc("/all-content", h.ViewAllContent).Queries(
		"page", "{page}", "count", "{count}").Methods(http.MethodGet)

	client.HandleFunc("/content/{content_id}", h.ViewConcreteContentByID).Methods(http.MethodGet)

	client.HandleFunc("/content-filter", h.GetContentByFilter).Methods(http.MethodGet)

	client.HandleFunc("/links/{content_id}", h.GetLinks).Methods(http.MethodGet)

	client.HandleFunc("/image/{content_id}", h.GetImage).Methods(http.MethodGet)

	// client's messages -------------------------------------------------------------------------------------------
	message := router.PathPrefix("/client/message").Subrouter()
	message.Use(h.ClientAuthentication)

	message.HandleFunc("", h.ViewAllMessages).Queries(
		"page", "{page}", "count", "{count}").Methods(http.MethodGet)

	message.HandleFunc("/unread", h.ViewUnreadMessages).Queries(
		"page", "{page}", "count", "{count}").Methods(http.MethodGet)

	message.HandleFunc("/read/{notification_id}", h.ReadMessage).Methods(http.MethodGet)

	message.HandleFunc("/mark/{notification_id}", h.MarkTheMessageAsRead).Methods(http.MethodPost)

	message.HandleFunc("/unmark/{notification_id}", h.UnmarkTheMessageAsRead).Methods(http.MethodPost)

	message.HandleFunc("/{notification_id}", h.DeleteMessageById).Methods(http.MethodDelete)

	message.HandleFunc("/mark-all", h.MarkAllMessagesAsRead).Methods(http.MethodPost)

	// Recommendations-------------------------------------------------------------------------------------------
	recommendations := router.PathPrefix("/client/recommendations").Subrouter()
	recommendations.Use(h.ClientAuthentication)

	recommendations.HandleFunc("/editorial", h.ViewAdminRecommendations).Queries(
		"page", "{page}", "count", "{count}").Methods(http.MethodGet)

	recommendations.HandleFunc("/user", h.ViewUserRecommendations).Queries(
		"page", "{page}", "count", "{count}").Methods(http.MethodGet)

	// Premieres -------------------------------------------------------------------------------------------
	premieres := router.PathPrefix("/client/premieres").Subrouter()
	premieres.HandleFunc("", h.ViewPremieres).Queries(
		"page", "{page}", "count", "{count}").Methods(http.MethodGet)

	//playlists -------------------------------------------------------------------------------------------
	playlist := router.PathPrefix("/client/playlist").Subrouter()
	playlist.Use(h.ClientAuthentication)

	playlist.HandleFunc("/{content_id}", h.AddToMyPlaylists).Queries(
		"playlist", "{playlist}").Methods(http.MethodPost)

	playlist.HandleFunc("", h.ViewFromMyPlaylist).Queries(
		"page", "{page}", "count", "{count}", "playlist", "{playlist}").Methods(http.MethodGet)

	playlist.HandleFunc("/{content_id}", h.DeleteContentFromPlaylist).Queries(
		"playlist", "{playlist}").Methods(http.MethodDelete)

	return router
}
