package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rkBekzat/films/internal/model"
)

type actor struct {
	db *sqlx.DB
}

func NewActor(db *sqlx.DB) Actor {
	return &actor{db: db}
}

func (a *actor) Create(user *model.Actor) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id", actorTable)
	row := a.db.QueryRow(query, user.Name, user.Gender, user.BirthDate)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (a *actor) Read(id string) (*model.Actor, error) {
	var res model.Actor
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", actorTable)
	err := a.db.Get(&res, query, id)
	return &res, err
}

func (a *actor) Update(user *model.Actor) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, gender=$2, birth_date=$3 WHERE id=$4", actorTable)
	_, err := a.db.Exec(query, user.Name, user.Gender, user.BirthDate, user.Id)
	return err
}

func (a *actor) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", actorTable)
	_, err := a.db.Exec(query, id)
	return err
}

func (a *actor) Search(text string) ([]model.Actor, error) {
	var res []model.Actor
	query := fmt.Sprintf("SELECT * FROM %s WHERE name LIKE $1", actorTable)
	err := a.db.Select(&res, query, text+"%")
	return res, err
}

func (a *actor) FilmedList(id string) ([]model.Film, error) {
	var res []model.Film
	query := fmt.Sprintf("SELECT title, description, rating FROM %s INNER JOIN %s WHERE %s.id = %s.film_id WHERE %s.actor_id == $1", filmTable, participantsTable, filmTable, participantsTable, participantsTable)
	err := a.db.Select(&res, query, id)
	if err != nil {
		return nil, err 
	}
	return res, nil 
}
