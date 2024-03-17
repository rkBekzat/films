package service

import "github.com/rkBekzat/films/internal/repository"

type account struct {
	repo repository.Account
}

func NewAccount(repo repository.Account) Account {
	return &account{
		repo: repo,
	}
}
