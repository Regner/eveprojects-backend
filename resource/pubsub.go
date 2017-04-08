package resource

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/regner/eveprojects-backend/esi"
	"github.com/regner/eveprojects-backend/model"
	"github.com/regner/eveprojects-backend/storage"
	"github.com/regner/eveprojects-backend/utils"
	"log"
	"net/http"
)

type PubSubResource struct {
	AllianceStorage    *storage.AllianceStorage
	CharacterStorage   *storage.CharacterStorage
	CorporationStorage *storage.CorporationStorage
}

type pushRequest struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

func (psr PubSubResource) UpdateCorporation(c *gin.Context) {
	var corporationID int
	var err error

	in := &pushRequest{}
	err = c.Bind(in)
	if err != nil {
		log.Printf("Error parsing PubSub push message: %v", err)
		c.String(http.StatusBadRequest, "")
		return
	}

	err = json.Unmarshal(in.Message.Data, &corporationID)
	if err != nil {
		log.Printf("Something went wrong unmarshling PubSub message: %v", err)
		c.String(http.StatusOK, "")
		return
	}

	corporation, err := psr.CorporationStorage.GetOne(corporationID)
	if err != nil && err == gorm.ErrRecordNotFound {
		corporation = model.Corporation{
			ID: corporationID,
		}

		err = psr.CorporationStorage.AddNew(corporation)
		if err != nil {
			c.String(http.StatusInternalServerError, "")
			return
		}
	} else if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	res, err := esi.GetCorporationInfo(corporationID)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	if res.StatusCode == http.StatusNotFound {
		log.Printf("Corporation ID not found: %d", corporationID)
		c.String(http.StatusOK, "")
		return
	}

	var esiCorporation esi.ESICorporation
	res.Body.FromJsonTo(&esiCorporation)

	corporation.Name = esiCorporation.Name
	corporation.Description = esiCorporation.Description
	corporation.Birthday = esiCorporation.Birthday
	corporation.Ticker = esiCorporation.Ticker
	corporation.URL = esiCorporation.URL
	corporation.MemberCount = esiCorporation.MemberCount
	corporation.TaxRate = esiCorporation.TaxRate
	corporation.AllianceID = esiCorporation.AllianceID
	corporation.CEOID = esiCorporation.CEOID
	corporation.FounderID = esiCorporation.FounderID

	err = psr.CorporationStorage.Update(corporation)
	if err != nil {

	}

	utils.UpdateAlliance(corporation.AllianceID)

	c.String(http.StatusOK, "OK")
}

func (psr PubSubResource) UpdateCharacter(c *gin.Context) {
	var characterID int
	var err error

	in := &pushRequest{}
	err = c.Bind(in)
	if err != nil {
		log.Printf("Error parsing PubSub push message: %v", err)
		c.String(http.StatusBadRequest, "")
		return
	}

	err = json.Unmarshal(in.Message.Data, &characterID)
	if err != nil {
		log.Printf("Something went wrong unmarshling PubSub message: %v", err)
		c.String(http.StatusBadRequest, "")
		return
	}

	character, err := psr.CharacterStorage.GetOne(characterID)
	if err != nil && err == gorm.ErrRecordNotFound {
		character = model.Character{
			ID: characterID,
		}

		err = psr.CharacterStorage.AddNew(character)
		if err != nil {
			c.String(http.StatusInternalServerError, "")
			return
		}
	} else if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	res, err := esi.GetCharacterInfo(characterID)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	if res.StatusCode == http.StatusNotFound {
		log.Printf("Character ID not found: %d", characterID)
		c.String(http.StatusOK, "")
		return
	}

	var esiCharacter esi.ESICharacter
	res.Body.FromJsonTo(&esiCharacter)

	character.Name = esiCharacter.Name
	character.Description = esiCharacter.Description
	character.Gender = esiCharacter.Gender
	character.Birthday = esiCharacter.Birthday
	character.CorporationID = esiCharacter.CorporationID
	character.AncestryID = esiCharacter.AncestryID
	character.BloodlineID = esiCharacter.BloodlineID
	character.RaceID = esiCharacter.RaceID
	character.SecurityStatus = esiCharacter.SecurityStatus

	err = psr.CharacterStorage.Update(character)
	if err != nil {

	}

	utils.UpdateCorporation(character.CorporationID)

	c.String(http.StatusOK, "OK")
}

func (psr PubSubResource) UpdateAlliance(c *gin.Context) {
}
