package main

import (
	"github.com/gin-gonic/gin"
	"github.com/regner/eveprojects-backend/resource"
	"github.com/regner/eveprojects-backend/storage"
	"github.com/regner/eveprojects-backend/utils"
)

func main() {
	r := gin.Default()

	utils.InitPubSub()

	db, err := storage.InitDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	characterStorage := storage.NewCharacterStorage(db)
	corporationStorage := storage.NewCorporationStorage(db)
	allianceStorage := storage.NewStorage(db)

	alliances := resource.AlliancesResource{
		AllianceStorage: allianceStorage,
	}

	characters := resource.CharactersResource{
		CharacterStorage: characterStorage,
	}

	corporations := resource.CorporationsResource{
		CorporationStorage: corporationStorage,
	}
	pubsub := resource.PubSubResource{
		AllianceStorage:    allianceStorage,
		CharacterStorage:   characterStorage,
		CorporationStorage: corporationStorage,
	}

	updates := resource.UpdateResource{
		AllianceStorage:    allianceStorage,
		CharacterStorage:   characterStorage,
		CorporationStorage: corporationStorage,
	}

	r.GET("/api/alliances/", alliances.GetAllAlliances)
	r.GET("/api/alliances/", alliances.AddAlliance)

	r.GET("/api/characters/", characters.GetAllCharacters)
	r.GET("/api/characters/", characters.AddCharacter)

	r.GET("/api/corporations/", corporations.GetAllCorporationss)
	r.GET("/api/corporations/", corporations.AddCorporation)

	r.GET("/api/pubsub/update_alliance/", pubsub.UpdateAlliance)
	r.GET("/api/pubsub/update_character/", pubsub.UpdateCharacter)
	r.GET("/api/pubsub/update_corporation/", pubsub.UpdateCorporation)

	r.GET("/api/update/alliances/", updates.UpdateAlliances)
	r.GET("/api/update/characters/", updates.UpdateCharacters)
	r.GET("/api/update/corporations/", updates.UpdateCorporations)
	r.GET("/api/update/inv_categories/", updates.UpdateInvCategories)
	r.GET("/api/update/inv_groups/", updates.UpdateInvGroups)
	r.GET("/api/update/inv_types/", updates.UpdateInvTypes)
	r.GET("/api/update/map_constellations/", updates.UpdateMapConstellations)
	r.GET("/api/update/map_regions/", updates.UpdateMapRegions)
	r.GET("/api/update/map_systems/", updates.UpdateMapSystems)

	r.Run(":8000")
}
