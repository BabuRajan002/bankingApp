package main

import (
	"bankingApp/app"
	"bankingApp/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
