package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rkBekzat/films/internal/model"
)

type Account interface {
	EmailExist(email string) (bool, error)
	CreateUser(user *model.User) error
	GetUser(username, password string) (*model.User, error)
}

type Actor interface {
	Create(user *model.Actor) (string, error)
	Read(id string) (*model.Actor, error)
	Update(user *model.Actor) error
	Delete(id string) error

	Search(text string) ([]model.Actor, error)
	FilmedList(id string) ([]model.Film, error)
}

type Film interface {
	Create(*model.Film) (string, error)
	GetById(id string) (*model.Film, error)
	GetFilms(offset, limit int, sortedBy, order string) ([]model.Film, error)
	Search(text string) ([]model.Film, error)
	Update(*model.Film) error
	Delete(id string) error
}

type Repo struct {
	AuthRepo  Account
	ActorRepo Actor
	FilmRepo  Film
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		AuthRepo:  NewAccount(db),
		ActorRepo: NewActor(db),
		FilmRepo:  NewFilm(db),
	}
}
