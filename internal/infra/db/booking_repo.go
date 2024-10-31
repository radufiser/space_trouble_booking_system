package db

import (
	"database/sql"

	"spacetrouble.com/booking/internal/domain"
)

type BookingRepositorySQL struct {
	DB *sql.DB
}

func NewBookingRepositorySQL(db *sql.DB) *BookingRepositorySQL {
	return &BookingRepositorySQL{DB: db}
}

func (repo *BookingRepositorySQL) FindAll() ([]*domain.Booking, error) {
	query := `SELECT id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date FROM bookings`

	rows, err := repo.DB.Query(query)
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
