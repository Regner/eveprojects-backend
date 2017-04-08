package model

import "time"

type MapConstellation struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
