package repository

import "github.com/jmoiron/sqlx"

type account struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) Account {
	return &account{db: db}
}
