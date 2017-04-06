package model

import "time"

type Corporation struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"ddescription" gorm:"size:2048"`
	Birthday    time.Time `json:"birthday"`
	Ticker      string    `json:"ticker"`
	URL         string    `json:"url"`
	MemberCount int       `json:"member_count"`
	TaxRate     float32   `json:"tax_rate"`
	AllianceID  int       `json:"alliance_id" gorm:"index"`

	CEO   Character
	CEOID int `json:"ceo_id"`

	Founder   Character
	FounderID int `json:"founder_id"`

	Characters []Character

	CreatedAt time.Time  `json:"ommit"`
	UpdatedAt time.Time  `json:"ommit"`
	DeletedAt *time.Time `json:"ommit"`
}
