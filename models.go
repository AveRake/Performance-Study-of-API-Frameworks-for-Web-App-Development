package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" validate:"required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}

type Recipe struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Ingredients string `json:"ingredients" validate:"required"`
	Instructions string `json:"instructions" validate:"required"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignkey:UserID"`
}
