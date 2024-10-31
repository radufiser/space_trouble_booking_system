package api

import (
	"encoding/json"
	"net/http"

	"spacetrouble.com/booking/internal/app"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
