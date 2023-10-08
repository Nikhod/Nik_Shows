package repositories

import (
	"Nik/pkg/models"
	"errors"
)

// Fills columns with information, else - return error
func (r *Repository) AddContentToDB(content *models.Content) error {
	var chekContent models.Content
	gorm := r.Db.Table("content").Where("name = ?", content.Name).First(&chekContent)
	affected := gorm.RowsAffected

	switch {
	case affected == 0:
		query := `insert into content (name, content_type_id, description, production_year,
                     age_limit, producers, directors, actors, main_characters,
                     duration, genre_id)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		err := r.Db.Exec(query, content.Name, content.ContentTypeID, content.Description,
			content.ProductionYear, content.AgeLimit, content.Producers, content.Directors, content.Actors,
			content.MainCharacters, content.Duration, content.GenreId).Error
		if err != nil {
			return err
		}

	case affected == 1 && chekContent.Active == false:
		query := `update content
 set active = true
 where name = ?;`
		err := r.Db.Exec(query, chekContent.Name).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// Fills columns with information, else - return error
func (r *Repository) AddLinksToDB(links *models.Links, contentID int) error {
	var chekLink models.Links
	gorm := r.Db.Table("links").Where("content_id = ?", contentID).First(&chekLink)

	affected := gorm.RowsAffected

	switch {
	case affected == 0:
		query := `insert into links (alfa, kinopoisk, okko, wink, content_id)
values (?,?,?,?,?);`
		err := r.Db.Exec(query, links.Alfa, links.Kinopoisk,
			links.Okko, links.Wink, contentID).Error
		if err != nil {
			return err
		}

	case affected == 1 && chekLink.Active == false:
		query := `update links
set active = true, updated_at = current_timestamp
where content_id = ?;`
		err := r.Db.Exec(query, contentID).Error
		if err != nil {
			return err
		}

	default:
		err := errors.New("Content updated!")
		return err
	}

	return nil
}

func (r *Repository) AddImageToDB(imageName string, contentID int) error {
	query := `update content
set cover_image_name = ?, updated_at = current_timestamp
where id = ?;`
	err := r.Db.Exec(query, imageName, contentID).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete Content in DB
func (r *Repository) DeleteContent(contentID int) error {
	query := `update content
set active = false, deleted_at = current_timestamp
where name = ?;`
	err := r.Db.Exec(query, contentID).Error
	if err != nil {
		return err
	}

	return nil
}

// Ban User using Login
func (r *Repository) BanUser(login string) error {
	query := `update users
set active = false, deleted_at = current_timestamp
where login = ?;`
	err := r.Db.Exec(query, login).Error
	if err != nil {
		return err
	}

	return nil
}

// Unblock User using Login
func (r *Repository) UnblockUser(login string) error {
	query := `update users
set active = true, updated_at = current_timestamp
where login = ?;`
	err := r.Db.Exec(query, login).Error
	if err != nil {
		return err
	}

	return nil
}

// Deactiveate link, else - return error
func (r *Repository) DeactivateLink(contentID int) error {
	query := `update links
set active = false
where content_id = ?;`
	err := r.Db.Exec(query, contentID).Error
	if err != nil {
		return err
	}

	return nil
}

// Add Content to recomendation DB, else - return error
func (r *Repository) AddRecommendationToDB(contentID int) error {
	if contentID == 0 {
		err := errors.New("There is no Content by this name!")
		return err
	}

	var recomend models.AdminRecommendations
	affected := r.Db.Table("admin_recommendations").Where("content_id = ?", contentID).First(&recomend).RowsAffected

	switch {
	case affected == 0:
		query := `insert into admin_recommendations(content_id)
values(?);`
		err := r.Db.Exec(query, contentID).Error
		if err != nil {
			return err
		}

	case affected == 1 && recomend.Active == false:
		query := `update admin_recommendations
set active = true, updated_at = current_timestamp
where id = ?;`
		err := r.Db.Exec(query, recomend.Id).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete recomendation from database, else - return error
func (r *Repository) DeleteRecommendationFromDB(contentID int) error {
	query := `update admin_recommendations
set active = false, deleted_at = current_timestamp
where content_id = ?;`
	err := r.Db.Exec(query, contentID).Error
	if err != nil {
		return err
	}

	return nil
}

// send message to EVERY user, else - return error
func (r *Repository) SendMessage(message *models.Notifications) error {
	user := make([]int, 10, 20)
	query := `select u.id
from users u;`
	gorm := r.Db.Raw(query).Scan(&user)
	err := gorm.Error
	if err != nil {
		return err
	}

	affected := gorm.RowsAffected

	query = `insert into notifications (notification, recipient_id)
values (?, ?);`

	// In the cycle only the recipients ID changes
	var recipientID int64
	for recipientID = 1; recipientID <= affected; recipientID++ {
		err = r.Db.Exec(query, message.Notification, recipientID).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// todo EditDescription
