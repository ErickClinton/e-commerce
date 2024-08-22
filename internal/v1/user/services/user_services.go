package services

import (
	"eccomerce/internal/v1/user/entity"
	"eccomerce/internal/v1/user/repository"
	"eccomerce/pkg/utils"
	"encoding/json"
)

type Service interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) Service {
	return &service{repo: repo}
}

func (s *service) Create(user *entity.User) error {
	userJSON, _ := json.MarshalIndent(user, "", "    ")
	utils.Logger.Info().Msgf("Start method create %v", string(userJSON))
	return s.repo.Create(user)
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

func (s *service) Update(user *entity.User) error {
	userJSON, _ := json.MarshalIndent(user, "", "    ")
	utils.Logger.Info().Msgf("Start method Update %v", string(userJSON))
	return s.repo.Update(user)
}

func (s *service) Delete(id uint) error {
	idJSON, _ := json.MarshalIndent(id, "", "    ")
	utils.Logger.Info().Msgf("Start method Delete %v", string(idJSON))
	return s.repo.Delete(id)
}
