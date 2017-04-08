package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/regner/eveprojects-backend/storage"
)

type UpdateResource struct {
	AllianceStorage    *storage.AllianceStorage
	CharacterStorage   *storage.CharacterStorage
	CorporationStorage *storage.CorporationStorage
}

func (psr UpdateResource) UpdateAlliances(c *gin.Context) {
}

func (psr UpdateResource) UpdateCharacters(c *gin.Context) {
}

func (psr UpdateResource) UpdateCorporations(c *gin.Context) {
}

func (psr UpdateResource) UpdateInvCategories(c *gin.Context) {
}

func (psr UpdateResource) UpdateInvGroups(c *gin.Context) {
}

func (psr UpdateResource) UpdateInvTypes(c *gin.Context) {
}

func (psr UpdateResource) UpdateMapConstellations(c *gin.Context) {
}

func (psr UpdateResource) UpdateMapRegions(c *gin.Context) {
}

func (psr UpdateResource) UpdateMapSystems(c *gin.Context) {
}
