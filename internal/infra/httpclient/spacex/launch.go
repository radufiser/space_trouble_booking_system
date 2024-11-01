package httpclient

import (
	"fmt"
	"time"

	"encoding/json"
	"net/http"
)

const LaunchesEndpoint = "/v4/launches/upcoming"

// Launch represents the fields of interest from each launch in the response.
type Launch struct {
	Launchpad    string `json:"launchpad"`
	Name         string `json:"name"`
	DateUnix     int64  `json:"date_unix"`
	FlightNumber int    `json:"flight_number"`
}

// LaunchClient handles requests to the SpaceX Launch API
type LaunchClient struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewLaunchClient initializes a new LaunchClient
func NewLaunchClient(baseURL string) *LaunchClient {
	return &LaunchClient{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    baseURL,
	}
}

// GetUpcomingLaunches fetches upcoming launches and extracts only the required fields
func (c *LaunchClient) GetUpcomingLaunches() ([]Launch, error) {
	req, err := http.NewRequest("GET", c.BaseURL+LaunchesEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var launches []Launch
	if err := json.NewDecoder(resp.Body).Decode(&launches); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return launches, nil
}
