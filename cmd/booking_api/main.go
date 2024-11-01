package main

import (
	"log"
	"net/http"

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
	launchpadClient := httpclient.NewLaunchpadClient("https://api.spacexdata.com")

	// Setup repositories
	bookingRepo := db.NewBookingRepository(database)
	launchpadRepo := db.NewLaunchpadRepository(database)
	destinationRepo := db.NewDestinationRepository(database)
	scheduleRepo := db.NewScheduleRepository(database)

	// Initialize services and handlers
	bookingService := app.NewBookingService(bookingRepo)
	bookingHandler := api.NewBookingHandler(bookingService)

	// Syncing
	launchpadService := app.NewLaunchpadService(launchpadRepo, launchpadClient)
	err = launchpadService.SyncActiveLaunchpads()
	if err != nil {
		log.Fatalf("Sync of active launchpads failed: %v", err)
	}

	scheduleService := app.NewScheduleService(destinationRepo, launchpadRepo, scheduleRepo)
	err = scheduleService.GenerateWeeklySchedule()
	if err != nil {
		log.Fatalf("Sync of active launchpads failed: %v", err)
	}

	// Setup router and start the server
	router := api.SetupRouter(bookingHandler)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
