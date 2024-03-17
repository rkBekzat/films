package service

import (
	"github.com/rkBekzat/films/internal/model"
	"github.com/rkBekzat/films/internal/repository"
)

type Account interface {
	CreateUser(user *model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (string, string, error)
}

type Actor interface {
	Create(user *model.Actor) (string, error)
	Read(id string) (*model.Actor, error)
	Update(user *model.Actor) error
	Delete(id string) error
}

type Service struct {
	AuthService  Account
	ActorService Actor
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		AuthService: NewAuthorization(repo.AuthRepo),
		ActorService: NewActor(repo.ActorRepo),
	}
}
