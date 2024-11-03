package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// Setup router
	router := api.SetupRouter(bookingHandler)

	// Create the server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Channel to listen for interrupt or terminate signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Channel to signal server shutdown is complete
	shutdownComplete := make(chan struct{})

	// Start server in a separate goroutine
	go func() {
		log.Println("Server running on port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for signal
	<-signalChan
	log.Println("Shutting down server...")

	// Create a context with a timeout for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	close(shutdownComplete)
	log.Println("Server gracefully stopped")

	<-shutdownComplete
}
