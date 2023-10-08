package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Database's Types
type (
	Content struct {
		Id             int       `json:"id"`
		Name           string    `json:"name"`
		ContentTypeID  int       `json:"content_type_id"`
		Description    string    `json:"description"`
		ProductionYear int       `json:"production_year"`
		AgeLimit       int       `json:"age_limit"`
		Producers      string    `json:"producers"`
		Directors      string    `json:"directors"`
		Actors         string    `json:"actors"`
		MainCharacters string    `json:"main_characters"`
		Duration       string    `json:"duration"`
		GenreId        int       `json:"genre_id"`
		ImageName      string    `json:"image_name"`
		Active         bool      `json:"active"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		DeletedAt      time.Time `json:"deleted_at"`
	}

	Links struct {
		Id        int    `json:"id"`
		Active    bool   `json:"active"`
		Alfa      string `json:"alfa"`
		Kinopoisk string `json:"kinopoisk"`
		Okko      string `json:"okko"`
		Wink      string `json:"wink"`
	}

	ContentType struct {
		ID   uint   `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	Genres struct {
		ID   uint   `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	Tokens struct {
		Id     int       `json:"id"`
		Token  string    `json:"token"`
		UserId int       `json:"user_id"`
		Expire time.Time `json:"expire"`
	}

	Users struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Age       int       `json:"age"`
		Login     string    `json:"login"`
		Password  string    `json:"password"`
		RoleId    int       `json:"role_id"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	Notifications struct {
		Id           int       `json:"id"`
		Notification string    `json:"notification"`
		RecipientId  int       `json:"recipient_id"`
		Active       bool      `json:"active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		DeletedAt    time.Time `json:"deleted_at"`
	}

	Roles struct {
		Id        int
		Role      string
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	Ratings struct {
		Id        int       `json:"id"`
		UserId    int       `json:"user_id"`
		ContentId int       `json:"content_id"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	WillWatch struct {
		Id        int       `json:"id"`
		UserId    int       `json:"user_id"`
		ContentId int       `json:"content_id"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	Favorites struct {
		Id        int       `json:"id"`
		UserId    int       `json:"user_id"`
		ContentId int       `json:"content_id"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	PlaylistTypes struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	Playists struct {
		Id         int       `json:"id"`
		PlaylistId int       `json:"playlist_id"`
		UserId     int       `json:"user_id"`
		ContentId  int       `json:"content_id"`
		Active     bool      `json:"active"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		DeletedAt  time.Time `json:"deleted_at"`
	}

	AdminRecommendations struct {
		Id        int       `json:"id"`
		ContentID int       `json:"content_id"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)

// types for Work
type (
	Answer struct {
		Date           time.Time
		ResponseAnswer string
	}

	SendToken struct {
		Date   time.Time
		Answer string
		Token  string
	}

	TokenClaims struct {
		jwt.StandardClaims
		Login string
	}

	ContentAndLinks struct {
		Content Content `json:"content"`
		Links   Links   `json:"links"`
	}

	Pagination struct {
		Count  int
		Offset int
	}

	PUC struct {
		PlaylistId int
		UserId     int
		ContentId  int
	}

	Filter struct {
		Genre  string `json:"genre"`
		Year   int    `json:"production_year"`
		Actors string `json:"actors"`
		Page   int    `json:"page"`
		Count  int    `json:"count"`
	}

	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	Db struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	}

	Configs struct {
		Server Server `json:"server"`
		Db     Db     `json:"db"`
	}

	Image struct {
		ContentID int `json:"content_id"`
	}
)

// types for Send
type (
	Support struct {
		SupportText string `json:"support_text"`
		Email       string `json:"email"`
	}

	SendNotification struct {
		Id           int       `json:"id"`
		Date         time.Time `json:"date"`
		Notification string    `json:"notification"`
	}

	SendContents struct {
		ContentID      int    `json:"content_id"`
		Name           string `json:"name"`
		ProductionYear int    `json:"production_year" gorm:"production_year" `
		Genre          string `json:"genre"`
		Actors         string `json:"actors"`
	}

	SendConcreteContent struct {
		ID             uint   `gorm:"column:content_id"`
		Name           string `gorm:"column:name"`
		ContentType    string `gorm:"column:content_type"`
		Description    string `gorm:"column:description"`
		ProductionYear int    `gorm:"column:production_year"`
		Producers      string `gorm:"column:producers"`
		Directors      string `gorm:"column:directors"`
		Actors         string `gorm:"column:actors"`
		MainCharacters string `gorm:"column:main_characters"`
		Duration       string `gorm:"column:duration"`
		Genre          string `gorm:"column:genre"`
	}

	SendLinks struct {
		Alfa      string `json:"alfa"`
		Kinopoisk string `json:"kinopoisk"`
		Okko      string `json:"okko"`
		Wink      string `json:"wink"`
	}
)
