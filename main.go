package main

import (
	"github.com/joho/godotenv"
	"github.com/nickypangers/banking/app"
	"github.com/nickypangers/banking/logger"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
		return
	}
	logger.Info("Starting application")
	app.Start()
}
