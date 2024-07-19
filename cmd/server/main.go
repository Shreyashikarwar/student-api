package main

import (
	"student-api/config"
	"student-api/routers"
)

func main() {
	config.ConnectDatabase()

	router := routers.SetupRouter()
	router.Run(":8080")
}
