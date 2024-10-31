package api

import (
	"github.com/gorilla/mux"
)

func SetupRouter(handler *BookingHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/bookings", handler.GetBookings).Methods("GET")
	return r
}
