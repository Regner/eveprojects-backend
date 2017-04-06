package esi

import (
	"fmt"
	"github.com/franela/goreq"
	"log"
	"time"
)

type ESICharacter struct {
	AncestryID     int       `json:"ancestry_id,omitempty"`
	Birthday       time.Time `json:"birthday,omitempty"`
	BloodlineID    int       `json:"bloodline_id,omitempty"`
	CorporationID  int       `json:"corporation_id,omitempty"`
	Description    string    `json:"description,omitempty"`
	Gender         string    `json:"gender,omitempty"`
	Name           string    `json:"name,omitempty"`
	RaceID         int       `json:"race_id,omitempty"`
	SecurityStatus float32   `json:"security_status,omitempty"`
}

func GetCharacterInfo(id int) (*goreq.Response, error) {
	res, err := goreq.Request{
		Method: "GET",
		Uri:    fmt.Sprintf("https://esi.tech.ccp.is/latest/characters/%d/", id),
	}.Do()

	if err != nil {
		log.Printf("Error while requesting character info from ESI: %v", err)
	}

	return res, err
}
