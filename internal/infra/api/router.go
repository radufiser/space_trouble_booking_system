package api

import (
	"github.com/gorilla/mux"
)

func SetupRouter(handler *BookingHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/bookings", handler.GetBookings).Methods("GET")
	r.HandleFunc("/bookings", handler.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", handler.DeleteBooking).Methods("DELETE")
	return r
}
