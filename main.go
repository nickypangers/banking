package main

import (
	"github.com/nickypangers/banking/app"
	"github.com/nickypangers/banking/logger"
)

func main() {
	logger.Info("Starting application")
	app.Start()
}
