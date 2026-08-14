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

func (s *CinemaService) GetAllCinema() ([]model.Cinema, error) {
	query := `
		SELECT id, name, location, rate
		FROM cinema
		ORDER BY id
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cinemas []model.Cinema

	for rows.Next() {
		var cinema model.Cinema

		err := rows.Scan(
			&cinema.ID,
			&cinema.Name,
			&cinema.Location,
			&cinema.Rate,
		)

		if err != nil {
			return nil, err
		}

		cinemas = append(cinemas, cinema)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cinemas, nil
}

func (s *CinemaService) GetCinemaById(id int) (model.Cinema, error) {
	var cinema model.Cinema

	query := `
		SELECT id, name, location, rate
		FROM cinema
		WHERE id = $1
	`

	err := s.db.QueryRow(
		query,
		id,
	).Scan(
		&cinema.ID,
		&cinema.Name,
		&cinema.Location,
		&cinema.Rate,
	)

	if err != nil {
		return cinema, err
	}

	return cinema, nil
}

func (s *CinemaService) UpdateCinema(id int, cinema model.Cinema) (*model.Cinema, error) {

	if cinema.Name == "" {
		return nil, errors.New("name should be filled")
	}

	if cinema.Location == "" {
		return nil, errors.New("location should be filled")
	}

	query := `
		UPDATE cinema
		SET name = $1,
		    location = $2,
		    rate = $3
		WHERE id = $4
		RETURNING id, name, location, rate
	`

	var updatedCinema model.Cinema

	err := s.db.QueryRow(
		query,
		cinema.Name,
		cinema.Location,
		cinema.Rate,
		id,
	).Scan(
		&updatedCinema.ID,
		&updatedCinema.Name,
		&updatedCinema.Location,
		&updatedCinema.Rate,
	)

	if err != nil {
		return nil, err
	}

	return &updatedCinema, nil
}

func (s *CinemaService) DeleteCinema(id int) error {

	query := `
		DELETE FROM cinema
		WHERE id = $1
	`

	result, err := s.db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
