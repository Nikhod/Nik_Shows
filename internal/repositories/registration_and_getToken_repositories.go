package repositories

import (
	"Nik/pkg/models"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

// return true if Login is free, else - false
func (r *Repository) IsLoginFree(login string) bool {
	var user models.Users
	amountOfChar := r.Db.Table("users").Where("login = ?", login).First(&user).RowsAffected
	if amountOfChar != 0 {
		return false
	}
	return true
}

// Adds user using name, login and hashed pass
func (r *Repository) AddUserToDb(user *models.Users, hash string) error {
	query := `insert into users (name, age, login, password)
values (?,?,?,?);`
	err := r.Db.Exec(query, user.Name, user.Age, user.Login, hash).Error
	if err != nil {
		return err
	}
	return nil
}

// Add token to DB
func (r *Repository) AddTokenToDb(userId int, token string) error {
	tokens := models.Tokens{}
	affected := r.Db.Table("tokens").Where("user_id = ?", userId).First(&tokens).RowsAffected

	if affected == 0 {
		query := `insert into tokens(token, user_id)
values(?, ?);`
		err := r.Db.Exec(query, token, userId).Error
		if err != nil {
			return err
		}

	} else {
		query := `update tokens
set token = ?, updated_at = ?
where tokens.user_id = ?;`
		err := r.Db.Exec(query, token, time.Now(), userId).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// Get user using token
func (r *Repository) GetUserByToken(token string) (userId int, err error) {
	query := `select user_id
from tokens
where token = ?;`
	err = r.Db.Raw(query, token).Scan(&userId).Error
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// Return role_id (1 - Admin ; 2 - Client)
func (r *Repository) IsAdmin(userId int) (roleId int, err error) {
	query := `select role_id
from users 
where id = ?;`
	err = r.Db.Raw(query, userId).Scan(&roleId).Error
	if err != nil {
		return 0, err
	}
	return roleId, nil
}

// Return active of user
func (r *Repository) GetUserActiveByID(userId int) (active bool, err error) {
	query := `select active
from users u
where u.id = ?;`
	err = r.Db.Raw(query, userId).Scan(&active).Error
	if err != nil {
		return false, err
	}

	return active, nil
}
