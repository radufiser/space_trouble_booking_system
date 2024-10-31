package api

import (
	"time"

	"spacetrouble.com/booking/internal/domain"
)

// BookingResponse represents the API response for a booking
type BookingResponse struct {
	ID            string    `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        *string   `json:"gender,omitempty"`
	Birthday      time.Time `json:"birthday"`
	LaunchpadID   string    `json:"launchpad_id"`
	DestinationID string    `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
}

// NewBookingResponse maps a domain.Booking to an API BookingResponse
func NewBookingResponse(booking *domain.Booking) BookingResponse {
	return BookingResponse{
		ID:            booking.ID,
		FirstName:     booking.FirstName,
		LastName:      booking.LastName,
		Gender:        booking.Gender,
		Birthday:      booking.Birthday,
		LaunchpadID:   booking.LaunchpadID,
		DestinationID: booking.DestinationID,
		LaunchDate:    booking.LaunchDate,
	}
}
