package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Track model
type Track struct {
	gorm.Model
	Name   string
	Start  *time.Time
	Stop   *time.Time
	UserID uint
}
