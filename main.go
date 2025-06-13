package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"thugcorp.io/nomado/data"
	"thugcorp.io/nomado/handlers"
	"thugcorp.io/nomado/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("nomado.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {

	//Initialize the logger
	logInstance := initializeLogger()

	// Environment setup
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Database connection setup
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize data repository
	propertyRepository, err := data.NewPropertyRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize property repository: %v", err)
		logInstance.Error("Failed to initialize property repository", err)
	}

	// Property handler initialization
	propertyHandler := handlers.PropertyHandler{}
	propertyHandler.Logger = logInstance
	propertyHandler.Storage = propertyRepository

	// Serve API endpoints
	http.HandleFunc("/api/v1/properties/top", propertyHandler.GetTopProperties)
	http.HandleFunc("/api/v1/properties/random", propertyHandler.GetRandomProperties)

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", addr, err)
		logInstance.Error("Failed to start server", err)
	}
}
