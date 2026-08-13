package service

import (
	"database/sql"
	"errors"

	"github.com/piipiets/go-cinema-restapi/model"
)

type CinemaService struct {
	db *sql.DB
}

func NewCinemaService(db *sql.DB) *CinemaService {
	return &CinemaService{
		db: db,
	}
}

func (s *CinemaService) CreateCinema(cinema model.Cinema) (model.Cinema, error) {

	// Validation
	if cinema.Name == "" {
		return cinema, errors.New("name should be filled")
	}

	if cinema.Location == "" {
		return cinema, errors.New("location should be filled")
	}

	query := `
		INSERT INTO cinema (name, location, rate)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := s.db.QueryRow(
		query,
		cinema.Name,
		cinema.Location,
		cinema.Rate,
	).Scan(&cinema.ID)

	if err != nil {
		return cinema, err
	}

	return cinema, nil
}
