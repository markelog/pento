package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email  string `gorm:"not null;"`
	Active bool   `gorm:"default:false"`
}
