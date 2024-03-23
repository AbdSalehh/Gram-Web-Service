package main

import (
	"MyGram/database"
	"MyGram/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()

	PORT := os.Getenv("PORT")
	r.Run(":" + PORT)
}