package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/regner/eveprojects-backend/model"
	"log"
)

type CharacterStorage struct {
	db *gorm.DB
}

func NewCharacterStorage(db *gorm.DB) *CharacterStorage {
	return &CharacterStorage{db}
}

func (s CharacterStorage) FindAll() ([]model.Character, error) {
	var characters []model.Character

	err := s.db.Find(&characters).Error
	if err != nil {
		log.Printf("Error while getting all characters from the DB: %v", err)
		return nil, err
	}

	return characters, nil
}

func (s CharacterStorage) GetOne(id int) (model.Character, error) {
	var character model.Character

	err := s.db.First(&character, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("The requested character was nto found in the DB: %v", id)
		} else {
			log.Printf("Error while checking DB for an existing character: %v", err)
		}

		return model.Character{}, err
	}

	return character, nil
}

func (s CharacterStorage) AddNew(character model.Character) error {
	err := s.db.Create(&character).Error
	if err != nil {
		log.Printf("Error while saving new character to the database: %v", err)
	}

	return err
}

func (s CharacterStorage) Update(character model.Character) error {
	err := s.db.Save(character).Error
	if err != nil {
		log.Printf("Error while saving character to the database: %v", err)
	}

	return err
}

func (s CharacterStorage) IsExisting(id int) (bool, error) {
	_, err := s.GetOne(id)

	if err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	} else {
		return false, err
	}
}
