package database

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/recipe-app/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "recipes.db")
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Recipe{})
}
