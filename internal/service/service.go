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
