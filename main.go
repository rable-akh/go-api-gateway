package main

import (
	"go-api-gateway/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: loading .env file")
	}

	routes.RunRoutes()
}
