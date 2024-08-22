package services

import (
	"eccomerce/internal/v1/user/models"
	"eccomerce/internal/v1/user/repository"
)

type Service interface {
	RegisterUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) Service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *service) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *service) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
