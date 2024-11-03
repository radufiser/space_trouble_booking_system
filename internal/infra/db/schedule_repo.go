package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"spacetrouble.com/booking/internal/domain"
)

type ScheduleRepository struct {
	DB *sql.DB
}

func NewScheduleRepository(db *sql.DB) *ScheduleRepository {
	return &ScheduleRepository{DB: db}
}

func (repo ScheduleRepository) FetchSchedule(ctx context.Context, launchpadID string, dayOfWeek int, destinationID string) (*domain.WeeklySchedule, error) {
	var schedule domain.WeeklySchedule
	err := repo.DB.QueryRowContext(ctx, `
        SELECT launchpad_id, day_of_week, destination_id, last_updated
        FROM weekly_schedule
        WHERE launchpad_id = $1 AND day_of_week = $2 AND destination_id = $3
    `, launchpadID, dayOfWeek, destinationID).Scan(
		&schedule.LaunchpadID, &schedule.DayOfWeek, &schedule.DestinationID, &schedule.LastUpdated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("launch from launchpad %s on %s with destination %s %w",
				launchpadID, time.Weekday(dayOfWeek).String(), destinationID, domain.ErrNotFound)
		}
		return nil, fmt.Errorf("%w: %s", domain.ErrInternal, err)
	}

	return &schedule, nil
}
