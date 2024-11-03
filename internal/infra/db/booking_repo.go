package db

import (
	"context"
	"database/sql"
	"fmt"

	"spacetrouble.com/booking/internal/domain"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (repo *BookingRepository) FindAll(ctx context.Context) ([]*domain.Booking, error) {
	query := `SELECT id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date FROM bookings`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*domain.Booking
	for rows.Next() {
		booking := &domain.Booking{}
		if err := rows.Scan(
			&booking.ID, &booking.FirstName, &booking.LastName,
			&booking.Gender, &booking.Birthday, &booking.LaunchpadID,
			&booking.DestinationID, &booking.LaunchDate,
		); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (repo *BookingRepository) Create(ctx context.Context, booking *domain.Booking) error {
	query := `INSERT INTO bookings (id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.DB.ExecContext(ctx, query, booking.ID, booking.FirstName, booking.LastName, booking.Gender,
		booking.Birthday, booking.LaunchpadID, booking.DestinationID, booking.LaunchDate)

	return err
}

func (repo *BookingRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM bookings WHERE id = $1`

	result, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete booking: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	return nil
}
