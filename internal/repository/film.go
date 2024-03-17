package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rkBekzat/films/internal/model"
)

type film struct {
	db *sqlx.DB
}

func NewFilm(db *sqlx.DB) Film {
	return &film{db: db}
}

func (f *film) Create(flm *model.Film) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING id", filmTable)
	row := f.db.QueryRow(query, flm.Title, flm.Description, flm.ReleasedDate, flm.Rating)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (f *film) GetById(id string) (*model.Film, error) {
	var res model.Film
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", filmTable)
	err := f.db.Get(&res, query, id)
	return &res, err
}

func (f *film) GetFilms(offset, limit int, sortedBy, order string) ([]model.Film, error) {
	var res []model.Film
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s %s LIMIT %d OFFSET %d", filmTable, sortedBy, order, limit, offset)
	err := f.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *film) Search(text string) ([]model.Film, error) {
	var res []model.Film
	query := fmt.Sprintf("SELECT * FROM %s WHERE title LIKE $1", filmTable)
	err := f.db.Select(&res, query, text+"%")
	return res, err
}

func (f *film) Update(flm *model.Film) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, released_date=$3, rating=$4 WHERE id=$5", filmTable)
	_, err := f.db.Exec(query, flm.Title, flm.Description, flm.ReleasedDate, flm.Rating, flm.Id)
	return err
}

func (f *film) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", filmTable)
	_, err := f.db.Exec(query, id)
	return err
}
