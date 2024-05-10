package main

import (
	"banking/app"
	"banking/logger"
	"log"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatalf("Environment variables SERVER_ADDRESS and SERVER_PORT variables are not defined...")
	}
}

func main() {

	sanityCheck()
	logger.Info("Starting the application")
	app.Start()
}
