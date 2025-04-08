package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title        string   `json:"title" gorm:"not null"`
	Description  string   `json:"description"`
	PrepTime     string   `json:"prepTime"`
	CookTime     string   `json:"cookTime"`
	Image        string   `json:"image"`
	Ingredients  []string `json:"ingredients" gorm:"serializer:json"`
	Instructions []string `json:"instructions" gorm:"serializer:json"`
}
