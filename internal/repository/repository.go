package repository

import "github.com/jmoiron/sqlx"

type Account interface {
}

type Repo struct {
	AuthRepo Account
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		AuthRepo: NewAccount(db),
	}
}
