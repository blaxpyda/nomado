package main

import (
	"log"
	"net/http"

	"thugcorp.io/nomado/handlers"
	"thugcorp.io/nomado/logger"
)

func initializeLogger() *logger.Logger{
	logInstance, err := logger.NewLogger("nomado.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}


func main(){

	//Initialize the logger
	logInstance := initializeLogger()

	propertyHandler := handlers.PropertyHandler{}

	// Serve API endpoints
	http.HandleFunc("/api/v1/properties/top", propertyHandler.GetTopProperties)


	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", addr, err)
		logInstance.Error("Failed to start server", err)
	}
}