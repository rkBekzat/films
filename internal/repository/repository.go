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

type Repo struct {
	AuthRepo Account
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		AuthRepo: NewAccount(db),
	}
}
