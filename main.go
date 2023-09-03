package main

import (
	"Api-Rest-Gin/database"
	"Api-Rest-Gin/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	database.DBConnect()

	routes.HandleRequest()
}
