package domain

import "time"

type WeeklySchedule struct {
	LaunchpadID   string    
	DayOfWeek     int      
	DestinationID string   
	LastUpdated   time.Time
}
