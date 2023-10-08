package repositories

import "Nik/pkg/models"

// Make SQL Query, Return Content Id, using name of content, else - error
func (r *Repository) GetContentIdByName(nameOfContent string) (contentID int, err error) {
	query := `select content.id
from content
where content.name = ? and active = true;`
	err = r.Db.Raw(query, nameOfContent).Scan(&contentID).Error
	if err != nil {
		return 0, err
	}
	return contentID, nil
}

// Make SQL Query, Return struct Content, using contentID, else - error
func (r *Repository) GetContentAgeLimitByID(contentID int) (contentAgeLimit int, err error) {
	query := `select c.age_limit
from content c
where c.id = ?;`

	err = r.Db.Raw(query, contentID).Scan(&contentAgeLimit).Error
	if err != nil {
		return 0, err
	}

	return contentAgeLimit, nil
}

// Make SQL Query, Return struct User, using login, else - error
func (r *Repository) GetUserByLogin(login string) (user *models.Users, err error) {
	query := `select * 
from users
where login = ? and active = true;`
	err = r.Db.Raw(query, login).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// return struct user by id, else - error
func (r *Repository) GetUserByID(userID int) (user *models.Users, err error) {
	query := `select *
from users
where id = ?;`
	err = r.Db.Raw(query, userID).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// return USER-NAME by id, else - error
func (r *Repository) GetUserNameByID(userId int) (userName string, err error) {
	query := `select name
from users
where id = ?;`
	err = r.Db.Raw(query, userId).Scan(&userName).Error
	if err != nil {
		return "", err
	}

	return userName, nil
}

func (r *Repository) GetUserAge(userID int) (userAge int, err error) {
	query := `select u.age
from users u
where id = ?;`
	err = r.Db.Raw(query, userID).Scan(&userAge).Error
	if err != nil {
		return 0, err
	}

	return userAge, nil
}

// todo VlidateContentID (проверка, Есть ли вообще такой контент в БД)
