package repositories

import (
	"Nik/pkg/models"
	"errors"
	"time"
)

func (r *Repository) DeleteMyAccount(userID int) error {
	query := `update users
set active = false, updated_at = current_timestamp
where id = ?;`

	err := r.Db.Table("users").Exec(query, userID).Error
	if err != nil {
		return err
	}

	return nil
}

// View All Messages by User Id, else - return error
func (r *Repository) ViewAllMessages(userId int, pagination *models.Pagination) (messages []*models.SendNotification, err error) {
	query := `select n.id, n.created_at as date, n.notification as notification
from notifications n
where n.recipient_id = ? and active = true
order by n.id
limit ? offset ?;`
	err = r.Db.Raw(query, userId, pagination.Count, pagination.Offset).Scan(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// View UNREAD Messages by User Id, else - return error
func (r *Repository) ViewUnreadMessages(userID int, pagination *models.Pagination) (messages []*models.SendNotification, err error) {
	query := `select n.id, n.created_at as date, n.notification
from notifications n
where n.recipient_id = ?
  and is_read = false
  and n.active = true
order by n.id
limit ? offset ?;`
	err = r.Db.Raw(query, userID, pagination.Count, pagination.Offset).Scan(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// delete Notification by id, else - return error
func (r *Repository) DeleteMessageById(userID, notificationID int) error {
	query := `update notifications
set active = false, updated_at = current_timestamp
where id = ? and recipient_id = ?;`

	err := r.Db.Exec(query, notificationID, userID).Error
	if err != nil {
		return err
	}

	return nil
}

// Get message using user_ID and notification_ID, else - return error
func (r *Repository) GetMessage(userID, notificationID int) (message *models.SendNotification, err error) {
	query := `select n.id, n.created_at as date, n.notification
from notifications n
where n.id = ?
  and n.recipient_id = ?;`

	err = r.Db.Raw(query, notificationID, userID).Scan(&message).Error
	if err != nil {
		return nil, err
	}

	return message, nil
}

// Marks all messages as read, else - return error
func (r *Repository) MarkAsReadAllMessages(userID int) error {
	query := `update notifications
set is_read = true, updated_at = current_timestamp
where recipient_id = ?;`

	err := r.Db.Exec(query, userID).Error
	if err != nil {
		return err
	}

	return nil
}

// Marks concrete message as read, else - return error
func (r *Repository) MarkTheMessageAsRead(userID, notificationID int) error {
	query := `update notifications
set is_read     = true,
    updated_at = current_timestamp
where id = ?
  and recipient_id = ?;`
	err := r.Db.Exec(query, notificationID, userID).Error
	if err != nil {
		return err
	}

	return nil
}

// Unmarks concrete message as read, else - return error
func (r *Repository) UnmarkTheMessageAsRead(userID, notificationID int) error {
	query := `update notifications
set is_read     = false,
    updated_at = current_timestamp
where id = ?
  and recipient_id = ?;`
	err := r.Db.Exec(query, notificationID, userID).Error
	if err != nil {
		return err
	}

	return nil
}

// Get struct of SendRecommendations, else - error
func (r *Repository) GetUserRecommendations(userAge int, pagination *models.Pagination) (recommendations []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join genres g on c.genre_id = g.id
where c.age_limit <= ? limit ? offset ?;`
	err = r.Db.Raw(query, userAge, pagination.Count, pagination.Offset).Scan(&recommendations).Error
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

// Get struct of SendRecommendations, else - error
func (r *Repository) GetAdminRecommendations(pagination *models.Pagination) (recommendations []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from admin_recommendations a
         join content c on a.content_id = c.id
         join genres g on c.genre_id = g.id
where a.active = true
order by a.id
limit ? offset ?;`

	err = r.Db.Raw(query, pagination.Count, pagination.Offset).Scan(&recommendations).Error
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

// Get struct of SendRecommendations, but we can rename variable to "premieres";) , else - error
func (r *Repository) GetPremieres(pagination *models.Pagination) (premieres []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join genres g on c.genre_id = g.id
where production_year = ?
  and c.active = true
order by c.id
limit ? offset ?;`

	err = r.Db.Table("content").Raw(query, time.Now().Year(), pagination.Count, pagination.Offset).Scan(&premieres).Error
	if err != nil {
		return nil, err
	}

	return premieres, nil
}

// Get struct of SendRecommendations, but we can rename variable to "ALLcontents" ;) , else - error
func (r *Repository) GetAllContents(pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join genres g on c.genre_id = g.id
Where c.active = true
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, pagination.Count, pagination.Offset).Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, err
}

// Add the content to Db, else - return err
func (r *Repository) AddToMyPlaylist(puc *models.PUC) error {
	var checkContent models.Playists
	affected := r.Db.Table("playlists").
		Where("playlist_id = ? and user_id = ? and content_id = ?",
			puc.PlaylistId, puc.UserId, puc.ContentId).
		First(&checkContent).RowsAffected

	switch {
	case affected == 0:
		query := `insert into playlists (playlist_id, user_id, content_id)
values (?,?,?);`

		err := r.Db.Exec(query, puc.PlaylistId, puc.UserId, puc.ContentId).Error
		if err != nil {
			return err
		}
	case affected == 1 && checkContent.Active == false:
		query := `update playlists p
set active = true
where p.playlist_id = ?
  and p.user_id = ?
  and p.content_id = ?;`

		err := r.Db.Exec(query, puc.PlaylistId, puc.UserId, puc.ContentId).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// Return struct of SendRecommendations, else return error
func (r *Repository) GetContentsFromPlaylist(puc *models.PUC, pagination *models.Pagination) (myMovies []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from playlists p
         join users u on p.user_id = u.id
         join content c on p.content_id = c.id
         join genres g on c.genre_id = g.id
where p.active = true
  and p.id = ?
  and u.id = ?
order by p.id
limit ? offset ?;`

	err = r.Db.Raw(query, puc.PlaylistId, puc.UserId,
		pagination.Count, pagination.Offset).Scan(&myMovies).Error
	if err != nil {
		return nil, err
	}

	return myMovies, nil
}

// Delete content from playlist, else - return error
func (r *Repository) DeleteContentFromPlaylist(puc *models.PUC) error {
	query := `update playlists
set active     = false,
    deleted_at = current_timestamp
where user_id = ?
  and playlist_id = ?
  and content_id = ?;`

	err := r.Db.Exec(query, puc.UserId, puc.PlaylistId, puc.ContentId).Error
	if err != nil {
		return err
	}

	return nil
}

// get structs SendConcreteContent by ID, else - return error
func (r *Repository) GetConcreteContentByID(contentID int) (content *models.SendConcreteContent, err error) {
	query := `select c.id              as content_id,
       c.name            as name,
       ct.name           as content_type,
       c.description     as description,
       c.production_year as production_year,
       c.producers       as producers,
       c.directors       as directors,
       c.actors          as actors,
       c.main_characters as main_characters,
       c.duration        as duration,
       g.name            as genre
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.id = ?
  and c.active = true;`

	err = r.Db.Table("content").
		Raw(query, contentID).
		Scan(&content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

// get struct Links By content iD , else - return error
func (r *Repository) GetLinks(contentID int) (links *models.SendLinks, err error) {
	query := `select l.alfa as alfa, l.kinopoisk as kinopoisk, l.okko as okko, l.wink as wink
from links l
where l.content_id = ? and l.active = true;`

	err = r.Db.Raw(query, contentID).Scan(&links).Error
	if err != nil {
		return nil, err
	}

	return links, nil
}

// get structs SendConcreteContent by Year, else - return error
func (r *Repository) GetContentsByYear(year int, pagination *models.Pagination) (content []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true and c.production_year = ?
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, year, pagination.Count, pagination.Offset).
		Scan(&content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

// get structs SendConcreteContent by Actors, else - return error
func (r *Repository) GetContentsByActors(actors string, pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true
  and c.actors = ?
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, actors, pagination.Count, pagination.Offset).
		Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// get structs SendConcreteContent by Genre, else - return error
func (r *Repository) GetContentsByGenre(genre string, pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true
  and g.name = ?
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, genre, pagination.Count, pagination.Offset).
		Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// get structs SendConcreteContent by actors and Genre, else - return error
func (r *Repository) GetContentsByActorsAndGenre(filter *models.Filter, pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true
  and c.actors = ?
  and g.name = ?
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, filter.Actors, filter.Genre, pagination.Count, pagination.Offset).
		Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// get structs SendConcreteContent by Year and Genre, else - return error
func (r *Repository) GetContentsByYearAndGenre(filter *models.Filter, pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true
  and c.production_year = ?
  and g.name = ?
order by c.id
limit ? offset ?;`

	err = r.Db.Raw(query, filter.Year, filter.Genre, pagination.Count, pagination.Offset).
		Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// get structs SendConcreteContent by Year and Actors, else - return error
func (r *Repository) GetContentsByYearAndActors(filter *models.Filter, pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true
  and c.production_year = ?
  and c.actors = ?
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, filter.Year, filter.Actors, pagination.Count, pagination.Offset).
		Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// get structs SendConcreteContent by Filter, else - return error
func (r *Repository) GetContentsByFilters(filter *models.Filter, pagination *models.Pagination) (contents []*models.SendContents, err error) {
	query := `select c.id as content_id, c.name as name, c.production_year as production_year, g.name as genre, c.actors as actors
from content c
         join content_type ct on c.content_type_id = ct.id
         join genres g on c.genre_id = g.id
where c.active = true
  and g.name = ?
  and c.actors = ?
  and c.production_year = ?
order by c.id
limit ? offset ?;`
	err = r.Db.Raw(query, filter.Genre, filter.Actors, filter.Year, pagination.Count, pagination.Offset).
		Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (r *Repository) GetContentImage(contentID int) (imageName string, err error) {
	var checkContent models.Content
	affected := r.Db.Table("content").
		Where("id = ?", contentID).
		First(&checkContent).RowsAffected

	if affected == 0 {
		return "", errors.New("invalid data")
	}

	query := `select c.cover_image_name
from content c
where c.id = ?;`
	err = r.Db.Raw(query, contentID).
		Scan(&imageName).Error
	if err != nil {
		return "", err
	}

	return imageName, nil
}
