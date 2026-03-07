package service

import (
	"go-pos-app/internal/model"
	"go-pos-app/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetUsers() ([]model.User, error){
	return s.Repo.GetUsers()
}

func (s *UserService) CreateUsers(user model.User) error {
	return s.Repo.CreateUsers(user)
}