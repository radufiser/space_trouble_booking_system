package domain

type ScheduleRepository interface {
	FetchSchedule(launchpadID string, dayOfWeek int, destinationID string) (*WeeklySchedule, error)
}
