package main

import (
	"log"
	"net/http"
	"time"

	"spacetrouble.com/booking/config"
	"spacetrouble.com/booking/internal/app"
	"spacetrouble.com/booking/internal/infra/api"
	"spacetrouble.com/booking/internal/infra/db"
	httpclient "spacetrouble.com/booking/internal/infra/httpclient/spacex"
	"spacetrouble.com/booking/pkg"
)

func main() {
	cfg := config.LoadConfig()
	database, err := pkg.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize HTTP clients
	launchClient := httpclient.NewLaunchClient("https://api.spacexdata.com", 10*time.Second, 3*time.Hour)

	// Setup repositories
	bookingRepo := db.NewBookingRepository(database)
	scheduleRepo := db.NewScheduleRepository(database)
	destinationRepo := db.NewDestinationRepository(database)
	launchpadRepo := db.NewLaunchpadRepository(database)

	// Initialize services and handlers
	bookingService := app.NewBookingService(bookingRepo, scheduleRepo, destinationRepo, launchpadRepo, launchClient)
	bookingHandler := api.NewBookingHandler(bookingService)

	// Setup router and start the server
	router := api.SetupRouter(bookingHandler)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
