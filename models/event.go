package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserId      int       `json:"UserId"`
	User        User      `gorm:"foreignkey:UserId" json:"-"` //relatio ke table user
	Datetime    time.Time `json:"datetime" binding:"required"`
}
