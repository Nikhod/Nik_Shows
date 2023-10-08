package database

import (
	"Nik/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(config *models.Configs) (*gorm.DB, error) {
	dbUri := "host=" + config.Db.Host + " user=" + config.Db.User +
		" password=" + config.Db.Password + " dbname=" + config.Db.Database + " port=" + config.Db.Port
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
