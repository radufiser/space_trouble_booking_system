package db

import (
	"database/sql"
	"fmt"

	"spacetrouble.com/booking/internal/domain"
)

type LaunchpadRepository struct {
	DB *sql.DB
}

func NewLaunchpadRepository(db *sql.DB) *LaunchpadRepository {
	return &LaunchpadRepository{DB: db}
}

// SaveLaunchpads saves the launchpads to the database
func (r LaunchpadRepository) SaveLaunchpads(launchpads []domain.Launchpad) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
        INSERT INTO launchpads (id, name, full_name, locality, region, status)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (id) DO UPDATE
        SET name = EXCLUDED.name,
            full_name = EXCLUDED.full_name,
            locality = EXCLUDED.locality,
            region = EXCLUDED.region,
            status = EXCLUDED.status
    `)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, lp := range launchpads {
		if _, err := stmt.Exec(lp.ID, lp.Name, lp.FullName, lp.Locality, lp.Region, lp.Status); err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	return tx.Commit()
}

// GetAllLaunchpads retrieves all launchpads from the database
func (r LaunchpadRepository) GetAllLaunchpads() ([]domain.Launchpad, error) {
	rows, err := r.DB.Query("SELECT id, name, full_name, locality, region, status FROM launchpads")
	if err != nil {
		return nil, fmt.Errorf("failed to query launchpads: %w", err)
	}
	defer rows.Close()

	var launchpads []domain.Launchpad
	for rows.Next() {
		var lp domain.Launchpad
		if err := rows.Scan(&lp.ID, &lp.Name, &lp.FullName, &lp.Locality, &lp.Region, &lp.Status); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		launchpads = append(launchpads, lp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return launchpads, nil
}
