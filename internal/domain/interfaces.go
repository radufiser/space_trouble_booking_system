package domain

import "context"

//go:generate mockgen -source=internal/domain/interfaces.go -destination=internal/domain/mocks/mock_interfaces.go -package=mocks
type BookingRepository interface {
	FindAll(ctx context.Context) ([]*Booking, error)
	Create(ctx context.Context, booking *Booking) error
	Delete(ctx context.Context, id string) error
}

type DestinationRepository interface {
	FetchAllDestinations(ctx context.Context) ([]Destination, error)
	GetByID(ctx context.Context, id string) (*Destination, error)
}

type LaunchpadRepository interface {
	GetByID(ctx context.Context, id string) (*Launchpad, error)
}

type ScheduleRepository interface {
	FetchSchedule(ctx context.Context, launchpadID string, dayOfWeek int, destinationID string) (*WeeklySchedule, error)
}

type LaunchClient interface {
	GetUpcomingLaunches(ctx context.Context) ([]Launch, error)
}
