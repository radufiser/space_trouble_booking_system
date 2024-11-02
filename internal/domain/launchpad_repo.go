package domain

type LaunchpadRepository interface {
	GetByID(id string) (*Launchpad, error)
}
