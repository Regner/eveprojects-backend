package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/regner/eveprojects-backend/model"
	"os"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_PARAMS"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&model.Character{},
		&model.Alliance{},
		&model.Corporation{},
	)

	return db, nil
}
