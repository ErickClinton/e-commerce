package services

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/user/dto"
	"eccomerce/internal/v1/user/repository"
	"eccomerce/pkg/authentication"
	"eccomerce/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Service interface {
	utils.Service[dto.CreateUserRequest, entity.User]
	GetByEmail(email string) (*entity.User, error)
	Login(request *dto.LoginUserRequest) (string, error)
	UpdateById(request *dto.UpdateUserRequest, id int) error
}

type service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) Service {
	return &service{repo: repo}
}

func (s *service) Create(user *dto.CreateUserRequest) error {
	userJSON, _ := json.MarshalIndent(user, "", "    ")
	utils.Logger.Info().Msgf("Start method create %v", string(userJSON))
	hashedPassword, err := authentication.HashPassword(user.Password)
	if err != nil {
		return errors.New("Unable to create password. Please try again later.")
	}
	entityUser := &entity.User{
		Email:    user.Email,
		Password: hashedPassword,
		Role:     user.Role,
		Username: user.Username,
	}
	return s.repo.Create(entityUser)
}

func (s *service) GetByID(id uint) (*entity.User, error) {
	idJSON, _ := json.MarshalIndent(id, "", "    ")
	utils.Logger.Info().Msgf("Start method GetByID %v", string(idJSON))
	return s.repo.GetByID(id)
}

func (s *service) GetByEmail(email string) (*entity.User, error) {
	emailJSON, _ := json.MarshalIndent(email, "", "    ")
	utils.Logger.Info().Msgf("Start method GetByEmail %v", string(emailJSON))
	return s.repo.GetByEmail(email)
}

func (s *service) Update(user *dto.CreateUserRequest) error {
	userJSON, _ := json.MarshalIndent(user, "", "    ")
	utils.Logger.Info().Msgf("Start method Update %v", string(userJSON))
	entityUser := &entity.User{
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Username: user.Username,
	}
	return s.repo.Update(entityUser)
}

func (s *service) UpdateById(updateUserDto *dto.UpdateUserRequest, id int) error {
	userJSON, error := json.MarshalIndent(updateUserDto, "", "    ")
	if error != nil {
		return error
	}
	utils.Logger.Info().Msgf("Start method Update %v", string(userJSON))

	return s.repo.UpdateById(updateUserDto, id)
}

func (s *service) Delete(id uint) error {
	idJSON, _ := json.MarshalIndent(id, "", "    ")
	utils.Logger.Info().Msgf("Start method Delete %v", string(idJSON))
	return s.repo.Delete(id)
}

func (s *service) Login(request *dto.LoginUserRequest) (string, error) {
	requestJSON, _ := json.MarshalIndent(request, "", "    ")
	utils.Logger.Info().Msgf("Start method Login %v", string(requestJSON))

	user, err := s.repo.GetByEmail(request.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !authentication.CheckPasswordHash(request.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}
	secretKey := os.Getenv("SECRET_KEY")
	tokenManager := authentication.NewTokenManager(secretKey, time.Hour*24)

	token, err := tokenManager.GenerateToken(fmt.Sprintf("%d", user.ID), user.Role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
