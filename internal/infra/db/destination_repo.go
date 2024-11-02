package db

import (
	"database/sql"
	"errors"
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
func (repo DestinationRepository) FetchAllDestinations() ([]domain.Destination, error) {
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

func (repo *DestinationRepository) GetByID(id string) (*domain.Destination, error) {
	var destination domain.Destination
	err := repo.DB.QueryRow("SELECT id, name FROM destinations WHERE id = $1", id).
		Scan(&destination.ID, &destination.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: destination with ID %s", domain.ErrNotFound, id)
		}
		return nil, fmt.Errorf("%w: %s", domain.ErrInternal, err)
	}
	return &destination, nil
}
