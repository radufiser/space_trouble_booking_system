package app

import (
	"fmt"
	"log"

	"spacetrouble.com/booking/internal/domain"
	httpclient "spacetrouble.com/booking/internal/infra/httpclient/spacex"
)

// LaunchpadService handles operations related to launchpads
type LaunchpadService struct {
	repository domain.LaunchpadRepository
	httpClient *httpclient.LaunchpadClient
}

// NewLaunchpadService creates a new instance of LaunchpadService
func NewLaunchpadService(repository domain.LaunchpadRepository, httpClient *httpclient.LaunchpadClient) *LaunchpadService {
	return &LaunchpadService{
		repository: repository,
		httpClient: httpClient,
	}
}

// SyncActiveLaunchpads fetches active launchpads using the httpClient and saves them using the repository
func (s *LaunchpadService) SyncActiveLaunchpads() error {
	launchpads, err := s.httpClient.GetActiveLaunchpads()
	if err != nil {
		return fmt.Errorf("failed to fetch active launchpads: %w", err)
	}

	if err := s.repository.SaveLaunchpads(launchpads.ToLaunchpads()); err != nil {
		return fmt.Errorf("failed to save launchpads to database: %w", err)
	}

	log.Println("Successfully synchronized active launchpads.")
	return nil
}
