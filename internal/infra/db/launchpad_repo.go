package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"spacetrouble.com/booking/internal/domain"
)

type LaunchpadRepository struct {
	DB *sql.DB
}

func NewLaunchpadRepository(db *sql.DB) *LaunchpadRepository {
	return &LaunchpadRepository{DB: db}
}

func (repo *LaunchpadRepository) GetByID(ctx context.Context, id string) (*domain.Launchpad, error) {
	var launchpad domain.Launchpad
	err := repo.DB.QueryRowContext(ctx, `
		SELECT id, name, full_name, locality, region, status 
		FROM launchpads 
		WHERE id = $1`, id).
		Scan(&launchpad.ID, &launchpad.Name, &launchpad.FullName, &launchpad.Locality, &launchpad.Region, &launchpad.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: launchpad with ID %s", domain.ErrNotFound, id)
		}
		return nil, fmt.Errorf("%w: %s", domain.ErrInternal, err)
	}
	return &launchpad, nil
}
