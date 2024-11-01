package db

import (
	"database/sql"
	"fmt"

	"spacetrouble.com/booking/internal/domain"
)

type DestinationRepository struct {
	DB *sql.DB
}

func NewDestinationRepository(db *sql.DB) *DestinationRepository {
	return &DestinationRepository{DB: db}
}

// FetchDestinations retrieves the list of destinations from the database
func (repo DestinationRepository) FetchDestinations() ([]domain.Destination, error) {
	rows, err := repo.DB.Query("SELECT id, name FROM destinations")
	if err != nil {
		return nil, fmt.Errorf("failed to query destinations: %w", err)
	}
	defer rows.Close()

	var destinations []domain.Destination
	for rows.Next() {
		var destination domain.Destination
		if err := rows.Scan(&destination.ID, &destination.Name); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		destinations = append(destinations, destination)
	}

	return destinations, nil
}
