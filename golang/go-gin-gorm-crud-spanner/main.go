// Package main is the entry point for the application.
package main

import (
	"github.com/gitish/polyglot_training/config"
	"github.com/gitish/polyglot_training/routes"
)

// main initializes and starts the web server.
func main() {
	// Connect to Spanner
	config.ConnectDatabase()

	// Ensure the connection is closed when the application exits
	defer config.DisconnectDatabase()

	// Set up routes and start the server
	router := routes.SetupRouter(config.DB) // Pass the gorm.DB
	router.Run(":8081")                     // Start the server on port 8080
}
