package model

import "time"

type Character struct {
	ID             int       `json:"id" gorm:"primary_key"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty" gorm:"size:4096"`
	Gender         string    `json:"gneder"`
	Birthday       time.Time `json:"birthday"`
	CorporationID  int       `json:"corporation_id" gotm:"index"`
	AncestryID     int       `json:"ancestry_id"`
	BloodlineID    int       `json:"bloodline_id"`
	RaceID         int       `json:"race_id"`
	SecurityStatus float32   `json:"security_status"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
