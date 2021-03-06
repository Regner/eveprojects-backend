package model

import "time"

type MapSystem struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
