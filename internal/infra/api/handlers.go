package api

import (
	"encoding/json"
	"net/http"

	"spacetrouble.com/booking/internal/app"
	"spacetrouble.com/booking/internal/domain"
)

type BookingHandler struct {
	Service *app.BookingService
}

func NewBookingHandler(service *app.BookingService) *BookingHandler {
	return &BookingHandler{Service: service}
}

func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.Service.GetBookings()
	if err != nil {
		http.Error(w, "Failed to retrieve bookings", http.StatusInternalServerError)
		return
	}

	// Map domain.Booking to BookingResponse
	var response []BookingResponse
	for _, booking := range bookings {
		response = append(response, NewBookingResponse(booking))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var bookingReq BookingRequest

	// Decode and validate
	if err := json.NewDecoder(r.Body).Decode(&bookingReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := bookingReq.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	booking := &domain.Booking{
		FirstName:     bookingReq.FirstName,
		LastName:      bookingReq.LastName,
		Birthday:      bookingReq.Birthday.ToTime(),
		LaunchpadID:   bookingReq.LaunchpadID,
		DestinationID: bookingReq.DestinationID,
		LaunchDate:    bookingReq.LaunchDate.ToTime(),
	}

	if bookingReq.Gender != nil {
		booking.Gender = bookingReq.Gender
	}

	if err := h.Service.CreateBooking(booking); err != nil {
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}
