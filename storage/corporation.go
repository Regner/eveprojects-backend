package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/regner/eveprojects-backend/model"
	"log"
)

type CorporationStorage struct {
	db *gorm.DB
}

func NewCorporationStorage(db *gorm.DB) *CorporationStorage {
	return &CorporationStorage{db}
}

func (s CorporationStorage) FindAll() ([]model.Corporation, error) {
	var corporations []model.Corporation

	err := s.db.Find(&corporations).Error
	if err != nil {
		log.Printf("Error while getting all corporations from the DB: %v", err)
		return nil, err
	}

	return corporations, nil
}

func (s CorporationStorage) GetOne(id int) (model.Corporation, error) {
	var corporation model.Corporation

	err := s.db.First(&corporation, id).Error
	if err != nil {
		log.Printf("Error while checking DB for an existing corporation: %v", err)
		return model.Corporation{}, err
	}

	return corporation, nil
}

func (s CorporationStorage) AddNew(corporation model.Corporation) error {
	err := s.db.Create(&corporation).Error
	if err != nil {
		log.Printf("Error while saving new corporation to the database: %v", err)
	}

	return err
}

func (s CorporationStorage) Update(corporation model.Corporation) error {
	err := s.db.Save(corporation).Error
	if err != nil {
		log.Printf("Error while saving corporation to the database: %v", err)
	}

	return err
}

func (s CorporationStorage) IsExisting(id int) (bool, error) {
	_, err := s.GetOne(id)

	if err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	} else {
		return false, err
	}
}
