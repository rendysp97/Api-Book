package main

import (
	"api-book/database/connection"
	"api-book/router"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.StartServer().Run(":" + port)
	connection.ConnectDB()

	router.StartServer().Run(":" + port)
}
