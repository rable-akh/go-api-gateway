package config

import (
	"log"
	"os"
)

func DB_URI() string {
	dbURI := os.Getenv("DATABASE_URL")
	if dbURI == "" {
		log.Fatal("Error: can't find database url")
	}
	return dbURI
}

// Auth Service URI
func AuthServiceURI() string {
	authURI := os.Getenv("AUTH_SERVICE")
	if authURI == "" {
		log.Fatal("Error: can't find service url")
	}
	return authURI
}
