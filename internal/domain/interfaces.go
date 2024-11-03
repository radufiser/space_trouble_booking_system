package domain

//go:generate mockgen -source=internal/domain/interfaces.go -destination=internal/domain/mocks/mock_repository.go -package=mocks
type BookingRepository interface {
	FindAll() ([]*Booking, error)
	Create(booking *Booking) error
}

type DestinationRepository interface {
	FetchAllDestinations() ([]Destination, error)
	GetByID(id string) (*Destination, error)
}

type LaunchpadRepository interface {
	GetByID(id string) (*Launchpad, error)
}

type ScheduleRepository interface {
	FetchSchedule(launchpadID string, dayOfWeek int, destinationID string) (*WeeklySchedule, error)
}

type LaunchClient interface {
	GetUpcomingLaunches() ([]Launch, error)
}
