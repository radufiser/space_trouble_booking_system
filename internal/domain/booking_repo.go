package domain

type BookingRepository interface {
	FindAll() ([]*Booking, error)
}
