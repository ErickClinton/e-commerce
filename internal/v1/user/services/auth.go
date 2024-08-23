package services

import (
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

type ServiceAuth interface {
	Login(request *dto.LoginUserRequest) (string, error)
}

type serviceAuth struct {
	repo repository.UserRepository
}

func NewServiceAuth(repo repository.UserRepository) ServiceAuth {
	return &serviceAuth{repo: repo}
}

func (s *serviceAuth) Login(request *dto.LoginUserRequest) (string, error) {
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
