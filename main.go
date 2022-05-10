package main

import (
	"github.com/joho/godotenv"
	"github.com/nickypangers/banking-lib/logger"
	"github.com/nickypangers/banking/app"
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
