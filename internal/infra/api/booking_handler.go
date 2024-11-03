package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
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
	bookings, err := h.Service.GetBookings(r.Context())
	if err != nil {
		writeJSONError(w, "Failed to retrieve bookings", http.StatusInternalServerError)
		return
	}

	response := []BookingResponse{}
	for _, booking := range bookings {
		response = append(response, NewBookingResponse(booking))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var bookingReq BookingRequest

	if err := json.NewDecoder(r.Body).Decode(&bookingReq); err != nil {
		writeJSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := bookingReq.Validate(); err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
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

	err := h.Service.CreateBooking(r.Context(), booking)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeJSONError(w, err.Error(), http.StatusNotFound)
			return
		}
		if errors.Is(err, domain.ErrConflict) {
			writeJSONError(w, err.Error(), http.StatusConflict)
			return
		}
		if errors.Is(err, domain.ErrInvalidInput) {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Service.DeleteBooking(r.Context(), id)
	if err != nil {
		writeJSONError(w, "Failed to delete booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
