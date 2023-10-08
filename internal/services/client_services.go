package services

import (
	"Nik/pkg/helper"
	"Nik/pkg/models"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Service) DeleteMyAccount(userID int) error {
	err := s.Repository.DeleteMyAccount(userID)
	if err != nil {
		return err
	}

	return nil
}

// View All Messages by User Id, else - return error
func (s *Service) ViewAllMessages(userId int, strPage, strCount string) (messages []*models.SendNotification, err error) {
	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}

	messages, err = s.Repository.ViewAllMessages(userId, &pagination)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// View UNREAD Messages by User Id, else - return error
func (s *Service) ViewUnreadMessages(userId int, strPage, strCount string) (messages []*models.SendNotification, err error) {
	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}

	messages, err = s.Repository.ViewUnreadMessages(userId, &pagination)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Delete Notification by ID
func (s *Service) DeleteMessageById(userID int, strNotificationID string) error {
	notificationID, err := strconv.Atoi(strNotificationID)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteMessageById(userID, notificationID)
	if err != nil {
		return err
	}

	return nil
}

// read message by user_id and str_notification_id
func (s *Service) ReadMessage(userID int, strNotificationID string) (message *models.SendNotification, err error) {
	notificationID, err := strconv.Atoi(strNotificationID)
	if err != nil {
		return nil, err
	}

	message, err = s.Repository.GetMessage(userID, notificationID)
	if err != nil {
		return nil, err
	}

	err = s.Repository.MarkTheMessageAsRead(userID, notificationID)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// Marks all messages as read, else - return error
func (s *Service) MarkAsReadAllMessages(userID int) error {
	err := s.Repository.MarkAsReadAllMessages(userID)
	if err != nil {
		return err
	}

	return nil
}

// Mark concrete message, else - return error
func (s *Service) MarkTheMessageAsRead(userID int, strNotificationID string) error {
	notificationID, err := strconv.Atoi(strNotificationID)
	if err != nil {
		return err
	}

	err = s.Repository.MarkTheMessageAsRead(userID, notificationID)
	if err != nil {
		return err
	}

	return nil
}

// Mark concrete message, else - return error
func (s *Service) UnmarkTheMessageAsRead(userID int, strNotificationID string) error {
	notificationID, err := strconv.Atoi(strNotificationID)
	if err != nil {
		return err
	}

	err = s.Repository.UnmarkTheMessageAsRead(userID, notificationID)
	if err != nil {
		return err
	}

	return nil
}

// return struct with user's recommendations, else - error
func (s *Service) ViewUserRecommendations(userID int, strPage, strCount string) (recommendations []*models.SendContents, err error) {
	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}

	user, err := s.Repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	userAge := user.Age
	recommendations, err = s.Repository.GetUserRecommendations(userAge, &pagination)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

// return struct with admin's recommendations, else - error
func (s *Service) ViewAdminRecommendations(strPage, strCount string) (recommendations []*models.SendContents, err error) {
	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}

	recommendations, err = s.Repository.GetAdminRecommendations(&pagination)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

// return struct of SendContents, but we can rename variable "premieres" ;) , else - error
func (s *Service) ViewPremieres(strPage, strCount string) (premieres []*models.SendContents, err error) {
	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}

	premieres, err = s.Repository.GetPremieres(&pagination)
	if err != nil {
		return nil, err
	}

	return premieres, nil
}

// return struct of SendContents, but we can rename variable allContents ;) , else - error
func (s *Service) ViewAllContent(strPage, strCount string) (contents []*models.SendContents, err error) {
	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}

	contents, err = s.Repository.GetAllContents(&pagination)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// Add Content to one playlist, else - return error
func (s *Service) AddToMyPlaylist(userID int, strContentID, playlist string) error {
	playlistID := map[string]int{"my_movies": 1, "my_series": 2, "my_cartoons": 3,
		"will_watch": 4, "favorites": 5}

	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	puc := models.PUC{
		PlaylistId: playlistID[playlist],
		UserId:     userID,
		ContentId:  contentID,
	}

	err = s.Repository.AddToMyPlaylist(&puc)
	if err != nil {
		return err
	}

	return nil
}

// Return struct of SendContents, else return error
func (s *Service) ViewContentsFromMyPlaylist(userID int, strPage, strCount, playlist string) (contents []*models.SendContents, err error) {
	playlistID := map[string]int{"my_movies": 1, "my_series": 2, "my_cartoons": 3,
		"will_watch": 4, "favorites": 5}

	page, count, err := helper.ConvertToIntTheParams(strPage, strCount)
	if err != nil {
		return nil, err
	}

	pagination := models.Pagination{
		Count:  count,
		Offset: (page - 1) * count,
	}
	puc := models.PUC{
		PlaylistId: playlistID[playlist],
		UserId:     userID,
		ContentId:  0,
	}

	contents, err = s.Repository.GetContentsFromPlaylist(&puc, &pagination)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// delete content from one playlist, elsse - return error
func (s *Service) DeleteContentFromPlaylist(userID int, strContentID, playlist string) error {
	playlistID := map[string]int{"my_movies": 1, "my_series": 2, "my_cartoons": 3,
		"will_watch": 4, "favorites": 5}

	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	puc := models.PUC{
		PlaylistId: playlistID[playlist],
		UserId:     userID,
		ContentId:  contentID,
	}

	err = s.Repository.DeleteContentFromPlaylist(&puc)
	if err != nil {
		return err
	}

	return nil
}

// get struct SendConcreteContent by  id, else - return error
func (s *Service) ViewConcreteContentByID(userID, contentID int) (content *models.SendConcreteContent, err error) {
	userAge, err := s.Repository.GetUserAge(userID)
	if err != nil {
		return nil, err
	}

	contentAgeLimit, err := s.Repository.GetContentAgeLimitByID(contentID)
	if err != nil {
		return nil, err
	}

	if contentAgeLimit > userAge {
		err = errors.New("You haven't reached a certain age yet!")
		return nil, err
	}

	content, err = s.Repository.GetConcreteContentByID(contentID)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// get struct links content ID, else - return error
func (s *Service) GetLinks(contentID int) (links *models.SendLinks, err error) {
	links, err = s.Repository.GetLinks(contentID)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// get struct SendConcreteContent by filter, else - return error
func (s *Service) GetContentsByFilter(filter *models.Filter) (contents []*models.SendContents, err error) {
	pagination := models.Pagination{
		Count:  filter.Count,
		Offset: (filter.Page - 1) * filter.Count,
	}
	switch {
	case filter.Genre != "" && filter.Actors != "" && filter.Year != 0:
		contents, err = s.Repository.GetContentsByFilters(filter, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil
	case filter.Year != 0 && filter.Actors != "":
		contents, err = s.Repository.GetContentsByYearAndActors(filter, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil
	case filter.Year != 0 && filter.Genre != "":
		contents, err = s.Repository.GetContentsByYearAndGenre(filter, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil
	case filter.Actors != "" && filter.Genre != "":
		contents, err = s.Repository.GetContentsByActorsAndGenre(filter, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil
	case filter.Genre != "":
		contents, err = s.Repository.GetContentsByGenre(filter.Genre, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil
	case filter.Actors != "":
		contents, err = s.Repository.GetContentsByActors(filter.Actors, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil
	case filter.Year != 0:
		contents, err = s.Repository.GetContentsByYear(filter.Year, &pagination)
		if err != nil {
			return nil, err
		}

		return contents, nil

	default:
		err = errors.New("Your request does not meet the conditions!")
		return nil, err
	}
}

func (s *Service) GetImage(strContentID string) (image []byte, err error) {
	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return nil, err
	}

	imageName, err := s.Repository.GetContentImage(contentID)
	if err != nil {
		return nil, err
	}

	path := filepath.Join(imagesDirPath, imageName)

	image, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return image, nil
}
