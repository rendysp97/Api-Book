package main

import (
	"api-book/database/connection"
	"api-book/router"
	"os"
)

func main() {

	connection.ConnectDB()

	router.StartServer().Run(":" + os.Getenv("PORT"))
}
