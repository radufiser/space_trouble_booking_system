package main

import (
	"log"
	"net/http"

	"spacetrouble.com/booking/config"
	"spacetrouble.com/booking/internal/app"
	"spacetrouble.com/booking/internal/infra/api"
	"spacetrouble.com/booking/internal/infra/db"
	"spacetrouble.com/booking/pkg"
)

func main() {
	cfg := config.LoadConfig()
	database, err := pkg.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup repository, service, and handler
	bookingRepo := db.NewBookingRepositorySQL(database)
	bookingService := app.NewBookingService(bookingRepo)
	bookingHandler := api.NewBookingHandler(bookingService)

	// Setup router and start the server
	router := api.SetupRouter(bookingHandler)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
