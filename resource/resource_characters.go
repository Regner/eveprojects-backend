package resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/regner/eveprojects-backend/storage"
	"github.com/regner/eveprojects-backend/utils"
	"net/http"
)

type CharactersResource struct {
	CharacterStorage   *storage.CharacterStorage
	CorporationStorage *storage.CorporationStorage
	AllianceStorage    *storage.AllianceStorage
}

type addCharacter struct {
	CharacterID int `json:"character_id"`
}

func (cr CharactersResource) GetAllCharacters(c *gin.Context) {
	characters, err := cr.CharacterStorage.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (cr CharactersResource) AddCharacter(c *gin.Context) {
	in := &addCharacter{}

	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	existing, err := cr.CharacterStorage.IsExistingCharacter(in.CharacterID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	} else if existing {
		c.String(http.StatusBadRequest, fmt.Sprintf("Character %d already exists", in.CharacterID))
		return
	}

	utils.UpdateCharacter(in.CharacterID)

	c.JSON(http.StatusCreated, in)
}
