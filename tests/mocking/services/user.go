package services

import (
	"log"

	"github.com/rAndrade360/go-mock-tests/repositories"
)

type service struct {
	repo repositories.UserRepository
}

type UserService interface {
	Create(name, email string) error
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(name, email string) error {
	err := s.repo.Create(name, email)
	if err != nil {
		log.Println("Deu erro", err.Error())
	}

	return err
}
