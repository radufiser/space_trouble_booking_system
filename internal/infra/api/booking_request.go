package api

import (
	"errors"
	"time"
)

type Date time.Time

// UnmarshalJSON for Date parses "YYYY-MM-DD" format.
func (d *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Trim the double quotes

	// Parse date with custom format
	parsed, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}

	*d = Date(parsed)
	return nil
}

// To convert back to `time.Time` for database/storage purposes.
func (d Date) ToTime() time.Time {
	return time.Time(d)
}

type BookingRequest struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Gender        *string `json:"gender,omitempty"`
	Birthday      Date    `json:"birthday"`
	LaunchpadID   string  `json:"launchpad_id"`
	DestinationID string  `json:"destination_id"`
	LaunchDate    Date    `json:"launch_date"`
}

// Validate checks the required fields
func (b *BookingRequest) Validate() error {
	if b.FirstName == "" {
		return errors.New("first_name is required")
	}
	if b.LastName == "" {
		return errors.New("last_name is required")
	}
	if time.Time(b.Birthday).IsZero() {
		return errors.New("birthday is required")
	}
	if b.LaunchpadID == "" {
		return errors.New("launchpad_id is required")
	}
	if b.DestinationID == "" {
		return errors.New("destination_id is required")
	}
	if time.Time(b.LaunchDate).IsZero() {
		return errors.New("launch_date is required")
	}
	return nil
}
