package model

import "time"

type Alliance struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Ticker   string    `json:"ticker"`
	Birthday time.Time `json:"birthday"`

	ExecutorCorporation   Corporation
	ExecutorCorporationID int

	Corporations []Corporation

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
