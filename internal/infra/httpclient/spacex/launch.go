package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const LaunchesEndpoint = "/v4/launches/upcoming"

// Launch represents the fields of interest from each launch in the response.
type Launch struct {
	Launchpad    string    `json:"launchpad"`
	Name         string    `json:"name"`
	DateUTC      time.Time `json:"date_utc"`
	FlightNumber int       `json:"flight_number"`
}

// cacheEntry represents a cached item with an expiration time.
type cacheEntry struct {
	data      []Launch
	expiresAt time.Time
}

// LaunchClient handles requests to the SpaceX Launch API and manages a cache.
type LaunchClient struct {
	HTTPClient *http.Client
	BaseURL    string
	cache      cacheEntry    // Single cache entry for upcoming launches
	cacheTTL   time.Duration // Time-to-live for cache entries
}

func NewLaunchClient(baseURL string, httpTTL time.Duration, cacheTTL time.Duration) *LaunchClient {
	return &LaunchClient{
		HTTPClient: &http.Client{Timeout: httpTTL},
		BaseURL:    baseURL,
		cacheTTL:   cacheTTL,
	}
}

// GetUpcomingLaunches fetches upcoming launches and extracts only the required fields
func (c *LaunchClient) GetUpcomingLaunches() ([]Launch, error) {
	// Check if the result is in the cache and not expired
	if c.cache.data != nil && time.Now().Before(c.cache.expiresAt) {
		return c.cache.data, nil
	}

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

	// Store the result in the cache with an expiration time
	c.cache = cacheEntry{
		data:      launches,
		expiresAt: time.Now().Add(c.cacheTTL),
	}

	return launches, nil
}
