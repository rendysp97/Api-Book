package main

import (
	"api-book/database/connection"
	"api-book/router"
	"os"
)

func main() {
	var PORT = "8080"

	connection.ConnectDB()

	router.StartServer().Run(":" + os.Getenv(PORT))
}
