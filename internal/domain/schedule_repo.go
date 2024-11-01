package domain

type ScheduleRepository interface {
	SaveWeeklySchedule(schedule map[string][]string) error
}
