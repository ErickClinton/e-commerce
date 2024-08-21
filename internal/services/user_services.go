package services

import (
	"eccomerce/internal/models"
	"eccomerce/internal/repository"
	"eccomerce/internal/utils"
	"errors"
	"log"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(username, email, password, role string) error {
	println("Chegou!")
	if role != "user" && role != "admin" && role != "manager" {
		return errors.New("invalid role")
	}
	hash, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	user := models.User{
		Username: username,
		Email:    email,
		Password: hash,
		Role:     role,
	}

	return s.userRepo.CreateUser(user)
}
