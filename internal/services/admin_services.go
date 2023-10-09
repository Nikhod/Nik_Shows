package services

import (
	"Nik/pkg/models"
	"errors"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var (
	ErrInvalidData  = errors.New("invalid data")
	ErrRegistration = errors.New("incorrect iame/login/password entered for registration!")
	ErrLoginUsed    = errors.New("login is taken by another user")
)

const (
	imagesDirPath = "C:\\Users\\Lenovo\\projects\\MyShows_Analog\\Images"
)

func (s *Service) SaveImage(file io.Reader, fileName string) (string, error) {
	imageName := fileName[len(fileName)-5:]
	extension := uuid.New().String()

	path := filepath.Join(imagesDirPath, extension+imageName)

	imageFile, err := os.Create(path)
	if err != nil {
		return "", err
	}

	defer func(imageFile *os.File) {
		err := imageFile.Close()
		if err != nil {
		}
	}(imageFile)

	_, err = io.Copy(imageFile, file)
	if err != nil {
		return "", err
	}

	finalImageName := extension + imageName
	return finalImageName, nil
}

func (s *Service) ValidateImage(size int64) error {
	if size > 5_000_000_000 {
		return ErrInvalidData
	}
	return nil
}

func (s *Service) ValidateImageName(name string) error {
	if len(name) > 65 || len(name) < 5 {
		return ErrInvalidData
	}

	return nil
}

func (s *Service) AddImageToDB(imageName string, strContentID string) error {
	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	err = s.Repository.AddImageToDB(imageName, contentID)
	if err != nil {
		return err
	}

	return err
}

// Add content, else return error
func (s *Service) AddContent(content *models.Content, contentType string, genre string) error {
	//prepare the struct content to send to func
	typeOfContent := make(map[string]int, 3)
	typeOfContent["movie"], typeOfContent["series"], typeOfContent["cartoon"] = 1, 2, 3

	genreOfContent := make(map[string]int, 13)
	genreOfContent["fantasy"], genreOfContent["fantastic"], genreOfContent["thriller"],
		genreOfContent["horror"], genreOfContent["sport"], genreOfContent["adventure"],
		genreOfContent["crime"], genreOfContent["сomedy"], genreOfContent["biography"],
		genreOfContent["story"], genreOfContent["drama"], genreOfContent["detective"],
		genreOfContent["action_movie"] = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13

	content.ContentTypeID = typeOfContent[contentType]
	content.GenreId = genreOfContent[genre]

	// Sending struct to DB
	err := s.Repository.AddContentToDB(content)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddLinks(links *models.Links, nameOfContent string) error {
	contentID, err := s.Repository.GetContentIdByName(nameOfContent)
	if err != nil {
		return err
	}

	err = s.Repository.AddLinksToDB(links, contentID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateContentAndLinks(content *models.Content, links *models.Links) error {
	if len(content.Name) > 100 || len(content.Name) < 1 {
		return ErrInvalidData
	}
	if len(content.Description) > 300 || len(content.Description) < 5 {
		return ErrInvalidData
	}

	return nil
}

// Delete content in Db
func (s *Service) DeleteContent(strContentID string) error {
	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteContent(contentID)
	if err != nil {
		return err
	}
	return nil
}

// Ban user using the login
func (s *Service) BanUser(login string) error {
	err := s.Repository.BanUser(login)
	if err != nil {
		return err
	}

	return nil
}

// Unblock user Using login, else - return error
func (s *Service) UnblockUser(login string) error {
	err := s.Repository.UnblockUser(login)
	if err != nil {
		return err
	}
	return nil
}

// Deactiveate link, else - return error
func (s *Service) DeactivateLink(strContentID string) error {
	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	err = s.Repository.DeactivateLink(contentID)
	if err != nil {
		return err
	}

	return nil
}

// Add Content to recomendation DB, else - return error
func (s *Service) AddRecomendation(strContentID string) error {
	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	err = s.Repository.AddRecommendationToDB(contentID)
	if err != nil {
		return err
	}

	return nil
}

// Delete Content from recomendation, else - return error
func (s *Service) DeleteFromRecomendation(strContentID string) error {
	contentID, err := strconv.Atoi(strContentID)
	if err != nil {
		return err
	}

	err = s.Repository.DeleteRecommendationFromDB(contentID)
	if err != nil {
		return err
	}

	return nil
}

// send message to EVERY user, else - return error
func (s *Service) SendMessage(message *models.Notifications) error {
	err := s.Repository.SendMessage(message)
	if err != nil {
		return err
	}

	return nil
}
