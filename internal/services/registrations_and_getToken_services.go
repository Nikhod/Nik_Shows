package services

import (
	"Nik/internal/repositories"
	"Nik/pkg/helper"
	"Nik/pkg/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	admin = 1
)

type Service struct {
	Repository *repositories.Repository
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{Repository: repository}
}

// return true if Data for registration is correct, Else - false
func (s *Service) IsValidDataForRegistration(user *models.Users) (err error) {
	if len(user.Name) > 20 || len(user.Name) < 3 {
		return ErrRegistration
	}
	if len(user.Login) > 20 || len(user.Login) < 3 {
		return ErrRegistration
	}
	if len(user.Password) > 20 || len(user.Password) < 6 {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "_") || strings.Contains(user.Password, "-") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "@") || strings.Contains(user.Password, "#") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "$") || strings.Contains(user.Password, "%") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "&") || strings.Contains(user.Password, "*") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "(") || strings.Contains(user.Password, ")") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, ":") || strings.Contains(user.Password, ".") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "/") || strings.Contains(user.Password, `\`) {
		return ErrRegistration
	}
	if strings.Contains(user.Password, ",") || strings.Contains(user.Password, ";") {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "?") || strings.Contains(user.Password, `"`) {
		return ErrRegistration
	}
	if strings.Contains(user.Password, "!") || strings.Contains(user.Password, "~") {
		return ErrRegistration
	}
	return nil
}

// registers a user using login and password
func (s *Service) RegistrationUser(user *models.Users) (err error) {
	isFree := s.Repository.IsLoginFree(user.Login)
	if isFree == false {
		return ErrLoginUsed
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.Repository.AddUserToDb(user, string(hash))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GenerateToken(login, password string) (token string, err error) {
	user, err := s.Repository.GetUserByLogin(login)
	if err != nil {
		return "", err
	}

	userActive, err := s.GetActiveUserById(user.Id)
	if err != nil {
		return "", err
	}

	if userActive == false {
		err = errors.New("User is Blocked!")
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err = helper.CreateToken(user.Login)

	err = s.Repository.AddTokenToDb(user.Id, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Get User Using token
func (s *Service) GetUserIdByToken(token string) (userId int, err error) {
	userId, err = s.Repository.GetUserByToken(token)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// Return true if it is Admin, else false
func (s *Service) IsAdmin(userId int) (bool, error) {
	active, err := s.Repository.GetUserActiveByID(userId)
	if err != nil {
		return false, err
	}

	if active == false {
		err = errors.New("User doesn't exsist!")
		return false, err
	}

	isAdmin, err := s.Repository.IsAdmin(userId)
	if err != nil {
		return false, err
	}

	if isAdmin == admin {
		return true, nil
	} else {
		err = errors.New("Access denied!")
		return false, err
	}
}

// Return Active of user
func (s *Service) GetActiveUserById(userId int) (active bool, err error) {
	active, err = s.Repository.GetUserActiveByID(userId)
	if err != nil {
		return false, err
	}

	return active, nil
}
