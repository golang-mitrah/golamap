package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ClientID     string
	ClientSecret string
	TokenURL     string
	PlacesURL    string
)

func LoadConfig() {
	log.Println("Loading environment variables from .env file")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	TokenURL = os.Getenv("TOKEN_URL")
	PlacesURL = os.Getenv("PLACES_URL")
}
