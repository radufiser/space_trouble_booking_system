package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"spacetrouble.com/booking/internal/domain"
)

type BookingService struct {
	BookingRepo     domain.BookingRepository
	ScheduleRepo    domain.ScheduleRepository
	DestinationRepo domain.DestinationRepository
	LaunchpadRepo   domain.LaunchpadRepository
	LaunchClient    domain.LaunchClient
}

func NewBookingService(
	bookingRepo domain.BookingRepository,
	scheduleRepo domain.ScheduleRepository,
	destinationRepo domain.DestinationRepository,
	launchpadRepo domain.LaunchpadRepository,
	launchClient domain.LaunchClient) *BookingService {
	return &BookingService{
		BookingRepo:     bookingRepo,
		ScheduleRepo:    scheduleRepo,
		DestinationRepo: destinationRepo,
		LaunchpadRepo:   launchpadRepo,
		LaunchClient:    launchClient,
	}
}

func (s *BookingService) GetBookings(ctx context.Context) ([]*domain.Booking, error) {
	return s.BookingRepo.FindAll(ctx)
}

func (s *BookingService) DeleteBooking(ctx context.Context, id string) error {
	return s.BookingRepo.Delete(ctx, id)
}

func (s *BookingService) CreateBooking(ctx context.Context, booking *domain.Booking) error {
	if err := s.validateDestination(ctx, booking.DestinationID); err != nil {
		return err
	}

	if err := s.validateLaunchpad(ctx, booking.LaunchpadID); err != nil {
		return err
	}

	if err := s.checkFlightSchedule(ctx, booking); err != nil {
		return err
	}

	if err := s.checkLaunchConflicts(ctx, booking); err != nil {
		return err
	}

	booking.ID = uuid.New().String()
	return s.BookingRepo.Create(ctx, booking)
}

// validateDestination checks if the given destination ID is valid.
func (s *BookingService) validateDestination(ctx context.Context, destinationID string) error {
	_, err := s.DestinationRepo.GetByID(ctx, destinationID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("invalid destination ID: %w", err)
		}
		return fmt.Errorf("failed to validate destination ID: %w", domain.ErrInternal)
	}
	return nil
}

// validateLaunchpad checks if the given launchpad ID is valid.
func (s *BookingService) validateLaunchpad(ctx context.Context, launchpadID string) error {
	_, err := s.LaunchpadRepo.GetByID(ctx, launchpadID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("invalid launchpad ID: %w", err)
		}
		return fmt.Errorf("failed to validate launchpad ID: %w", domain.ErrInternal)
	}
	return nil
}

// checkFlightSchedule ensures that a flight is scheduled for the given booking details.
func (s *BookingService) checkFlightSchedule(ctx context.Context, booking *domain.Booking) error {
	_, err := s.ScheduleRepo.FetchSchedule(ctx,
		booking.LaunchpadID,
		int(booking.LaunchDate.Weekday()),
		booking.DestinationID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("flight not scheduled for this date and destination: %w", err)
		}
		return fmt.Errorf("failed to fetch schedule: %w", domain.ErrInternal)
	}
	return nil
}

// checkLaunchConflicts checks if there is a launch conflict for the given booking date.
func (s *BookingService) checkLaunchConflicts(ctx context.Context, booking *domain.Booking) error {
	launches, err := s.LaunchClient.GetUpcomingLaunches(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch upcoming launches: %w", err)
	}

	if findMatchingLaunch(launches, *booking) != nil {
		return fmt.Errorf("booking conflict with an existing launch: %w", domain.ErrConflict)
	}
	return nil
}

// findMatchingLaunch checks for any launch matching the given booking details.
func findMatchingLaunch(launches []domain.Launch, booking domain.Booking) *domain.Launch {
	for _, launch := range launches {
		if launch.LaunchpadId == booking.LaunchpadID && isSameDay(booking.LaunchDate, launch.Date) {
			return &launch
		}
	}
	return nil
}

// isSameDay checks if two times are on the same calendar day.
func isSameDay(time1, time2 time.Time) bool {
	year1, month1, day1 := time1.Date()
	year2, month2, day2 := time2.Date()
	return year1 == year2 && month1 == month2 && day1 == day2
}
