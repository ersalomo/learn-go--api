package service

import (
	"go-api/db/model"
	"go-api/db/repository"
)

type Service interface {
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	FindByEmail(email string) (model.User, error)
	Create(user model.User) (model.User, error)
}

type service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]model.User, error) {
	return s.repo.FindAll()
}
func (s *service) FindById(id int) (model.User, error) {
	return s.repo.FindById(id)
}
func (s *service) FindByEmail(email string) (model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *service) Create(user model.User) (model.User, error) {
	return s.repo.Create(user)
}
