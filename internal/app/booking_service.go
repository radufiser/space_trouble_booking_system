package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"spacetrouble.com/booking/internal/domain"
	spacex "spacetrouble.com/booking/internal/infra/httpclient/spacex"
)

type BookingService struct {
	BookingRepo     domain.BookingRepository
	ScheduleRepo    domain.ScheduleRepository
	DestinationRepo domain.DestinationRepository
	LaunchpadRepo   domain.LaunchpadRepository
	LaunchClient    *spacex.LaunchClient
}

func NewBookingService(
	bookingRepo domain.BookingRepository,
	scheduleRepo domain.ScheduleRepository,
	destinationRepo domain.DestinationRepository,
	launchpadRepo domain.LaunchpadRepository,
	launchClient *spacex.LaunchClient) *BookingService {
	return &BookingService{
		BookingRepo:     bookingRepo,
		ScheduleRepo:    scheduleRepo,
		DestinationRepo: destinationRepo,
		LaunchpadRepo:   launchpadRepo,
		LaunchClient:    launchClient}
}

func (s *BookingService) GetBookings() ([]*domain.Booking, error) {
	return s.BookingRepo.FindAll()
}

func (s *BookingService) CreateBooking(booking *domain.Booking) error {
	// Validate destination ID
	_, err := s.DestinationRepo.GetByID(booking.DestinationID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("booking creation failed: %w", err)
		}
		return fmt.Errorf("booking creation failed: %w", domain.ErrInternal)
	}

	// Validate launchpad ID
	_, err = s.LaunchpadRepo.GetByID(booking.LaunchpadID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("booking creation failed: %w", err)
		}
		return fmt.Errorf("booking creation failed: %w", domain.ErrInternal)
	}

	// Fetch schedule
	_, err = s.ScheduleRepo.FetchSchedule(
		booking.LaunchpadID,
		int(booking.LaunchDate.Weekday()),
		booking.DestinationID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("booking creation failed: %w", err)
		}
		return fmt.Errorf("booking creation failed: %w", domain.ErrInternal)
	}

	// Check for launch conflicts
	launches, err := s.LaunchClient.GetUpcomingLaunches()
	if err != nil {
		return fmt.Errorf("failed to fetch upcoming launches: %w", err)
	}

	launch := findMatchingLaunch(launches, *booking)
	if launch != nil {
		return fmt.Errorf("booking not possible: %w", domain.ErrConflict)
	}

	// Create the booking
	booking.ID = uuid.New().String()

	return s.BookingRepo.Create(booking)
}

func findMatchingLaunch(launches []spacex.Launch, booking domain.Booking) *spacex.Launch {
	for _, launch := range launches {
		if launch.Launchpad == booking.LaunchpadID && isSameDay(booking.LaunchDate, launch.DateUTC) {
			return &launch
		}
	}

	return nil
}

func isSameDay(time1, time2 time.Time) bool {
	year1, month1, day1 := time1.Date()
	year2, month2, day2 := time2.Date()
	return year1 == year2 && month1 == month2 && day1 == day2
}
