package user

import (
	"github.com/jinzhu/gorm"

	"github.com/markelog/pento/back/database/models"
)

// User type
type User struct {
	db    *gorm.DB
	Model *models.User
}

// New User
func New(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

// StatusValue return value for Status method
type StatusValue struct {
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

// Status get user status
func (user *User) Status(email string) (*StatusValue, error) {
	var data models.User

	err := user.db.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}

	result := &StatusValue{
		Email:  data.Email,
		Active: data.Active,
	}

	return result, nil
}
