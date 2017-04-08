package resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/regner/eveprojects-backend/storage"
	"github.com/regner/eveprojects-backend/utils"
	"net/http"
)

type AlliancesResource struct {
	AllianceStorage *storage.AllianceStorage
}

type addAlliance struct {
	AllianceID int `json:"alliance_id"`
}

func (ar AlliancesResource) GetAllAlliances(c *gin.Context) {
	alliances, err := ar.AllianceStorage.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, alliances)
}

func (ar AlliancesResource) AddAlliance(c *gin.Context) {
	in := &addAlliance{}

	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	existing, err := ar.AllianceStorage.IsExisting(in.AllianceID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	} else if existing {
		c.String(http.StatusBadRequest, fmt.Sprintf("Alliance %d already exists", in.AllianceID))
		return
	}

	utils.UpdateAlliance(in.AllianceID)

	c.JSON(http.StatusCreated, in)
}
