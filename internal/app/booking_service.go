package app

import (
	"github.com/google/uuid"
	"spacetrouble.com/booking/internal/domain"
)

type BookingService struct {
	BookingRepo domain.BookingRepository
}

func NewBookingService(repo domain.BookingRepository) *BookingService {
	return &BookingService{BookingRepo: repo}
}

func (s *BookingService) GetBookings() ([]*domain.Booking, error) {
	return s.BookingRepo.FindAll()
}

func (s *BookingService) CreateBooking(booking *domain.Booking) error {
	booking.ID = uuid.New().String()

	//TODO Perform validation
	return s.BookingRepo.Create(booking)
}
