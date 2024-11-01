package httpclient

import (
	"bytes"
	"fmt"
	"time"

	"encoding/json"
	"net/http"

	"spacetrouble.com/booking/internal/domain"
)

const Endpoint = "/v4/launchpads/query"

// Launchpad represents a launchpad document in the response
type Launchpad struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Status   string `json:"status"`
	Locality string `json:"locality"`
	ID       string `json:"id"`
	Region   string `json:"region"`
}

// LaunchpadQueryRequest represents the JSON payload for the request
type LaunchpadQueryRequest struct {
	Query   map[string]interface{} `json:"query"`
	Options struct {
		Path   string         `json:"path"`
		Select map[string]int `json:"select"`
	} `json:"options"`
}

// LaunchpadQueryResponse represents the JSON response structure
type LaunchpadQueryResponse struct {
	Docs          []Launchpad `json:"docs"`
	TotalDocs     int         `json:"totalDocs"`
	Offset        int         `json:"offset"`
	Limit         int         `json:"limit"`
	TotalPages    int         `json:"totalPages"`
	Page          int         `json:"page"`
	PagingCounter int         `json:"pagingCounter"`
	HasPrevPage   bool        `json:"hasPrevPage"`
	HasNextPage   bool        `json:"hasNextPage"`
	PrevPage      *int        `json:"prevPage"`
	NextPage      *int        `json:"nextPage"`
}

func (r *LaunchpadQueryResponse) ToLaunchpads() []domain.Launchpad {
	var launchpads []domain.Launchpad
	for _, doc := range r.Docs {
		launchpad := domain.Launchpad{
			ID:       doc.ID,
			Name:     doc.Name,
			FullName: doc.FullName,
			Locality: doc.Locality,
			Region:   doc.Region,
			Status:   doc.Status,
		}
		launchpads = append(launchpads, launchpad)
	}
	return launchpads
}

type LaunchpadClient struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewLaunchpadClient(baseURL string) *LaunchpadClient {
	return &LaunchpadClient{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    baseURL,
	}
}

// GetActiveLaunchpads fetches launchpads that are not retired
func (c *LaunchpadClient) GetActiveLaunchpads() (*LaunchpadQueryResponse, error) {
	reqPayload := LaunchpadQueryRequest{
		Query: map[string]interface{}{
			"status": map[string]string{"$ne": "retired"},
		},
	}
	reqPayload.Options.Path = "docs"
	reqPayload.Options.Select = map[string]int{
		"name":      1,
		"full_name": 2,
		"status":    3,
		"locality":  4,
		"region":    5,
		"id":        6,
	}

	jsonData, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+Endpoint, bytes.NewBuffer(jsonData))
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

	var response LaunchpadQueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
