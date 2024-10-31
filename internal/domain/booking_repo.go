package domain

type BookingRepository interface {
	FindAll() ([]*Booking, error)
	Create(booking *Booking) error
}
