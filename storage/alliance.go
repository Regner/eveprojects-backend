package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/regner/eveprojects-backend/model"
)

type AllianceStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *AllianceStorage {
	return &AllianceStorage{db}
}

func (s AllianceStorage) All() ([]model.Alliance, error) {
	var alliances []model.Alliance

	err := s.db.Find(&alliances).Error
	if err != nil {
		return nil, err
	}

	return alliances, nil
}

func (s AllianceStorage) Get(id int) (model.Alliance, error) {
	var alliance model.Alliance

	err := s.db.First(&alliance, id).Error
	if err != nil {
		return model.Alliance{}, err
	}

	return alliance, nil
}

func (s AllianceStorage) Create(alliance model.Alliance) error {
	return s.db.Create(&alliance).Error
}

func (s AllianceStorage) IsExisting(id int) (bool, error) {
	_, err := s.Get(id)

	if err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	} else {
		return false, err
	}
}
