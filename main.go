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
	allianceStorage := storage.NewAllianceStorage(db)

	apiGroup := r.Group("/api/")
	{
		charactersGroup := apiGroup.Group("/characters/")
		{
			characters := resource.CharactersResource{
				characterStorage,
				corporationStorage,
				allianceStorage,
			}

			charactersGroup.GET("/", characters.GetAllCharacters)
			charactersGroup.POST("/", characters.AddCharacter)
		}

		corporationsGroup := apiGroup.Group("/corporations/")
		{
			corporations := resource.CorporationsResource{
				characterStorage,
				corporationStorage,
				allianceStorage,
			}

			corporationsGroup.GET("/", corporations.GetAllCorporationss)
			corporationsGroup.POST("/", corporations.AddCorporation)
		}

		alliancesGroup := apiGroup.Group("/alliances/")
		{
			alliances := resource.AlliancesResource{
				characterStorage,
				corporationStorage,
				allianceStorage,
			}

			alliancesGroup.GET("/", alliances.GetAllAlliances)
			alliancesGroup.POST("/", alliances.AddAlliance)
		}

		pubsubGroup := apiGroup.Group("/pubsub/")
		{
			pubsub := resource.PubSubResource{
				characterStorage,
				corporationStorage,
				allianceStorage,
			}

			pubsubGroup.POST("/update_corporation/", pubsub.UpdateCorporation)
			pubsubGroup.POST("/update_character/", pubsub.UpdateCharacter)
			pubsubGroup.POST("/update_alliance/", pubsub.UpdateAlliance)
		}

	}

	r.Run(":8000")
}
