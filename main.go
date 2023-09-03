package main

import (
	"Api-Rest-Gin/database"
	"Api-Rest-Gin/routes"
)

func main() {
	database.DBConnect()

	routes.HandleRequest()
}
