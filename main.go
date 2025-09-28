package main

import (
	"api-book/database/connection"
	"api-book/router"
)

func main() {
	var PORT = ":8080"

	connection.ConnectDB()

	router.StartServer().Run(PORT)
}
