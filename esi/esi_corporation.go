package esi

import (
	"fmt"
	"github.com/franela/goreq"
	"log"
	"time"
)

type ESICorporation struct {
	Name        string    `json:"corporation_name,omitempty"`
	Description string    `json:"corporation_description,omitempty"`
	Birthday    time.Time `json:"creation_date,omitempty"`
	Ticker      string    `json:"ticker,omitempty"`
	URL         string    `json:"url,omitempty"`
	MemberCount int       `json:"member_count,omitempty"`
	TaxRate     float32   `json:"tax_rate,omitempty"`
	AllianceID  int       `json:"alliance_id,omitempty"`
	CEOID       int       `json:"ceo_id,omitempty"`
	FounderID   int       `json:"creator_id,omitempty"`
}

func GetCorporationInfo(id int) (*goreq.Response, error) {
	res, err := goreq.Request{
		Method: "GET",
		Uri:    fmt.Sprintf("https://esi.tech.ccp.is/latest/corporations/%d/", id),
	}.Do()

	if err != nil {
		log.Printf("Error while requesting corporation info from ESI: %v", err)
	}

	return res, err
}
