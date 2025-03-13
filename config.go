package config

import (
	"os"
	"log"
)

var SecretKey string

func LoadConfig() {
	SecretKey = os.Getenv("SECRET_KEY")
	if SecretKey == "" {
		log.Fatal("SECRET_KEY environment variable is required")
	}
}
