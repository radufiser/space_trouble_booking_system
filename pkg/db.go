package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"spacetrouble.com/booking/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database successfully")
	return db, nil
}
