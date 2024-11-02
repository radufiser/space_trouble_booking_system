package domain

type DestinationRepository interface {
	FetchAllDestinations() ([]Destination, error)
	GetByID(id string) (*Destination, error)
}
