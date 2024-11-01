package app

import (
	"fmt"

	"spacetrouble.com/booking/internal/domain"
)

type ScheduleService struct {
	DestinationRepo domain.DestinationRepository
	LaunchpadRepo   domain.LaunchpadRepository
	ScheduleRepo    domain.ScheduleRepository
}

func NewScheduleService(
	destinationRepo domain.DestinationRepository,
	launchpadRepo domain.LaunchpadRepository,
	scheduleRepo domain.ScheduleRepository,
) *ScheduleService {
	return &ScheduleService{
		DestinationRepo: destinationRepo,
		LaunchpadRepo:   launchpadRepo,
		ScheduleRepo:    scheduleRepo,
	}
}

// GenerateWeeklySchedule creates a weekly schedule for each launchpad, assigning destinations and saves it to database
func (s ScheduleService) GenerateWeeklySchedule() error {
	destinations, err := s.DestinationRepo.FetchDestinations()
	if err != nil {
		return fmt.Errorf("error fetching destinations: %w", err)
	}

	launchpads, err := s.LaunchpadRepo.GetAllLaunchpads()
	if err != nil {
		return fmt.Errorf("error fetching launchpads: %w", err)
	}

	schedule := make(map[string][]string)

	for i, launchpad := range launchpads {
		weekSchedule := make([]string, 7)

		for day := 0; day < 7; day++ {
			destination := destinations[(i+day)%len(destinations)]
			weekSchedule[day] = destination.ID
		}
		schedule[launchpad.ID] = weekSchedule
	}

	err = s.ScheduleRepo.SaveWeeklySchedule(schedule)
	if err != nil {
		return fmt.Errorf("error saving weekly schedule: %w", err)
	}

	return nil
}
