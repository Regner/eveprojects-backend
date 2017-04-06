package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/regner/eveprojects-backend/model"
)

type AllianceStorage struct {
	db *gorm.DB
}

func NewAllianceStorage(db *gorm.DB) *AllianceStorage {
	return &AllianceStorage{db}
}

func (s AllianceStorage) FindAll() ([]model.Alliance, error) {
	var alliances []model.Alliance

	err := s.db.Find(&alliances).Error
	if err != nil {
		return nil, err
	}

	return alliances, nil
}

func (s AllianceStorage) GetOne(id int) (model.Alliance, error) {
	var alliance model.Alliance

	err := s.db.First(&alliance, id).Error
	if err != nil {
		return model.Alliance{}, err
	}

	return alliance, nil
}

func (s AllianceStorage) AddNewAlliance(alliance model.Alliance) error {
	return s.db.Create(&alliance).Error
}

func (s AllianceStorage) IsExistingAlliance(id int) (bool, error) {
	_, err := s.GetOne(id)

	if err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	} else {
		return false, err
	}
}
