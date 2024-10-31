package app

import "spacetrouble.com/booking/internal/domain"

type BookingService struct {
	BookingRepo domain.BookingRepository
}

func NewBookingService(repo domain.BookingRepository) *BookingService {
	return &BookingService{BookingRepo: repo}
}

func (s *BookingService) GetBookings() ([]*domain.Booking, error) {
	return s.BookingRepo.FindAll()
}
