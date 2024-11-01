package db

import (
	"database/sql"
	"fmt"
)

type ScheduleRepository struct {
	DB *sql.DB
}

func NewScheduleRepository(db *sql.DB) *ScheduleRepository {
	return &ScheduleRepository{DB: db}
}

func (repo ScheduleRepository) SaveWeeklySchedule(schedule map[string][]string) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
        INSERT INTO weekly_schedule (launchpad_id, day_of_week, destination_id, last_updated)
        VALUES ($1, $2, $3, NOW())
        ON CONFLICT (launchpad_id, day_of_week) DO UPDATE
        SET destination_id = EXCLUDED.destination_id, last_updated = EXCLUDED.last_updated
    `)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for launchpadID, days := range schedule {
		for dayOfWeek, destinationID := range days {
			if _, err := stmt.Exec(launchpadID, dayOfWeek, destinationID); err != nil {
				return fmt.Errorf("failed to execute statement: %w", err)
			}
		}
	}
	return tx.Commit()
}
