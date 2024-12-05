package config

import (
	"github.com/joho/godotenv"
	"log"
)

const TempPath = "./data/books/"

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
