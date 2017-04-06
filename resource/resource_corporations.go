package resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/regner/eveprojects-backend/storage"
	"github.com/regner/eveprojects-backend/utils"
	"net/http"
)

type CorporationsResource struct {
	CharacterStorage   *storage.CharacterStorage
	CorporationStorage *storage.CorporationStorage
	AllianceStorage    *storage.AllianceStorage
}

type addCorporation struct {
	CorporationID int `json:"corporation_id"`
}

func (cr CorporationsResource) GetAllCorporationss(c *gin.Context) {
	corporations, err := cr.CorporationStorage.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, corporations)
}

func (cr CorporationsResource) AddCorporation(c *gin.Context) {
	in := &addCorporation{}

	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	existing, err := cr.CorporationStorage.IsExistingCorporation(in.CorporationID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	} else if existing {
		c.String(http.StatusBadRequest, fmt.Sprintf("Corporation %d already exists", in.CorporationID))
		return
	}

	utils.UpdateCorporation(in.CorporationID)

	c.JSON(http.StatusCreated, in)
}
