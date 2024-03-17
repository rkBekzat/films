package service

import (
	"github.com/rkBekzat/films/internal/model"
	"github.com/rkBekzat/films/internal/repository"
)

type actor struct {
	repo repository.Actor
}

func NewActor(repo repository.Actor) Actor {
	return &actor{repo: repo}
}

func (a *actor) Create(user *model.Actor) (string, error) {
	return a.repo.Create(user)
}

func (a *actor) Read(id string) (*model.Actor, error) {
	return a.repo.Read(id)
}

func (a *actor) Update(user *model.Actor) error {
	return a.repo.Update(user)
}

func (a *actor) Delete(id string) error {
	return a.repo.Delete(id)
}

func (a *actor) Search(text string) ([]model.Actor, error) {
	return a.repo.Search(text)
}

func (a *actor) FilmedList(id string) ([]model.Film, error) {
	return a.FilmedList(id)
}
