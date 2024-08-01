package main

import (
	"hotel_booking/microservice/routes"
	"hotel_booking/microservice/utility"
	"net/http"
)

func main() {
	logger := utility.InitLogger()

	router := routes.SetupRoutes(logger)

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal("Server failed to start: %v", err)
	}
}
