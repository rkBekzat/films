package service

import "github.com/rkBekzat/films/internal/repository"

type Account interface {
}

type Service struct {
	AuthService Account
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		AuthService: NewAccount(repo.AuthRepo),
	}
}
