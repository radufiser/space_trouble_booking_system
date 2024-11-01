package domain

type LaunchpadRepository interface {
	SaveLaunchpads(launchpads []Launchpad) error
	GetAllLaunchpads() ([]Launchpad, error)
}
