package main

import (
	"backend/config"
)

func main() {
	router := config.SetupRouter()
	port := config.Port()

	router.Run(port)
}
