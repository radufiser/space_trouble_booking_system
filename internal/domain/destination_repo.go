package domain

type DestinationRepository interface {
	FetchDestinations() ([]Destination, error)
}
