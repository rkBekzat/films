package service

import (
	"github.com/rkBekzat/films/internal/model"
	"github.com/rkBekzat/films/internal/repository"
)

type Account interface {
	CreateUser(user *model.User) error 
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Service struct {
	AuthService Account
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		AuthService: NewAuthorization(repo.AuthRepo),
	}
}
