package service

import (
	"errors"
	"strconv"

	"github.com/rkBekzat/films/internal/model"
	"github.com/rkBekzat/films/internal/repository"
)

type film struct {
	repo repository.Film
}

func NewFilm(repo repository.Film) Film {
	return &film{repo: repo}
}

func (f *film) Create(fm *model.Film) (string, error) {
	return f.repo.Create(fm)
}

func (f *film) GetById(id string) (*model.Film, error) {
	return f.repo.GetById(id)
}

func (f *film) GetFilms(o, l, sortedBy, order string) ([]model.Film, error) {
	if order != "ASC" && order != "DESC" {
		return nil, errors.New("this order doesn't exist. Choose ASC or DESC")
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		return nil, errors.New("failed to parse offset")
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		return nil, errors.New("failed to parse limit")
	}

	return f.repo.GetFilms(offset, limit, sortedBy, order)
}

func (f *film) Search(text string) ([]model.Film, error) {
	return f.repo.Search(text)
}

func (f *film) Update(fm *model.Film) error {
	return f.repo.Update(fm)
}

func (f *film) Delete(id string) error {
	return f.repo.Delete(id)
}
