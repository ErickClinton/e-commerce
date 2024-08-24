package services

import (
	"eccomerce/internal/v1/auth/dto"
	"eccomerce/internal/v1/user/services"
	"eccomerce/pkg/authentication"
	"eccomerce/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ServiceAuth interface {
	Login(request *dto.LoginUserRequest) (string, error)
}

type serviceAuth struct {
	userService services.Service
}

func NewServiceAuth(userService services.Service) ServiceAuth {
	return &serviceAuth{userService: userService}
}

func (s *serviceAuth) Login(request *dto.LoginUserRequest) (string, error) {
	requestJSON, _ := json.MarshalIndent(request, "", "    ")
	utils.Logger.Info().Msgf("Start method Login %v", string(requestJSON))

	user, err := s.userService.GetByEmail(request.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !authentication.CheckPasswordHash(request.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY não definida nas variáveis de ambiente")
	}
	tokenManager := authentication.NewTokenManager(secretKey, time.Hour*24)

	token, err := tokenManager.GenerateToken(fmt.Sprintf("%d", user.ID), user.Role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
