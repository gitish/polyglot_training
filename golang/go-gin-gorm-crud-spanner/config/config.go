package config

import (
	"fmt"
	"log"

	"github.com/gitish/polyglot_training/models"
	spannergorm "github.com/googleapis/go-gorm-spanner"
	_ "github.com/googleapis/go-sql-spanner"
	"gorm.io/gorm"
)

var DB *gorm.DB // GORM DB // it is a wrapper around the underlying database connection pool which is managed by the database/sql package in Go

// ConnectDatabase initializes the 'DB' and connects to the database
func ConnectDatabase() {
	projectId, instanceId, databaseId := "your_projectId", "your_instanceId", "your_databaseId"
	dsn := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instanceId, databaseId)
	database, err := gorm.Open(spannergorm.New(spannergorm.Config{DriverName: "spanner", DSN: dsn}), &gorm.Config{
		// SkipDefaultTransaction: true, // Disable Default Transaction: GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency, you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}

	// Set the global Db variable
	DB = database

	// Check the connection
	var msg []string
	if err := DB.Raw("SELECT 'Hello World!'").Scan(&msg).Error; err != nil {
		log.Fatalf("failed to execute query: %v", err)
		return
	}
	if len(msg) == 0 {
		log.Fatalf("failed to execute query")
		return
	}
	for _, m := range msg {
		fmt.Printf("%s\n", m)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{}) // Automatically migrate your schema, to keep your schema up to date.
}

// DisconnectDatabase closes the connection with the Database server.
func DisconnectDatabase() {
	sqlDB, err := DB.DB() // access the underlying *sql.DB
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	if err := sqlDB.Close(); err != nil { // close the connection
		log.Fatalf("failed to close database connection: %v", err)
	}

	fmt.Println("Database connection closed successfully.")
}
