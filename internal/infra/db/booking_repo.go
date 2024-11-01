package db

import (
	"database/sql"

	"spacetrouble.com/booking/internal/domain"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (repo *BookingRepository) FindAll() ([]*domain.Booking, error) {
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

func (repo *BookingRepository) Create(booking *domain.Booking) error {
	query := `INSERT INTO bookings (id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.DB.Exec(query, booking.ID, booking.FirstName, booking.LastName, booking.Gender,
		booking.Birthday, booking.LaunchpadID, booking.DestinationID, booking.LaunchDate)

	return err
}
