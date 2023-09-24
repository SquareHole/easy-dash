package models

import "time"

type Result struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Description string    `json:"description" gorm:"not null"`
	LastStatus  string    `json:"last_status" gorm:"not null"`
	LastUpdate  time.Time `json:"last_update" gorm:"null"`
	NextUpdate  time.Time `json:"next_update" gorm:"null"`
}
