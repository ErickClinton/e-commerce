package services

import (
	"eccomerce/configs"
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/user/dto"
	"eccomerce/internal/v1/user/repository"
	walletdto "eccomerce/internal/v1/wallet/dto"
	walletservices "eccomerce/internal/v1/wallet/services"
	"eccomerce/pkg/authentication"
	"eccomerce/pkg/utils"
	"encoding/json"
	"errors"
)

type Service interface {
	utils.Service[dto.CreateUserRequest, entity.User]
	GetByEmail(email string) (*entity.User, error)
	UpdateById(request *dto.UpdateUserRequest, id int) error
}

type service struct {
	repo          repository.UserRepository
	walletService walletservices.WalletService
}

func NewService(repo repository.UserRepository, walletService walletservices.WalletService) Service {
	return &service{
		repo:          repo,
		walletService: walletService,
	}
}

func (s *service) Create(user *dto.CreateUserRequest) error {
	userJSON, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		return err
	}
	utils.Logger.Info().Msgf("Start method create %v", string(userJSON))

	hashedPassword, err := authentication.HashPassword(user.Password)
	if err != nil {
		return errors.New("unable to create password. Please try again later")
	}

	entityUser := &entity.User{
		Email:    user.Email,
		Password: hashedPassword,
		Role:     user.Role,
		Username: user.Username,
	}

	if err := s.repo.Create(entityUser); err != nil {
		return err
	}

	walletRequest := &walletdto.CreateWalletRequest{
		UserId:  entityUser.ID,
		Balance: 0,
	}

	return s.walletService.Create(walletRequest)
}

func (s *service) GetByID(id uint) (*entity.User, error) {
	utils.Logger.Info().Msgf("Start method GetByID %d", id)
	return s.repo.GetByID(id)
}

func (s *service) GetByEmail(email string) (*entity.User, error) {
	utils.Logger.Info().Msgf("Start method GetByEmail %s", email)
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
	userJSON, err := json.MarshalIndent(updateUserDto, "", "    ")
	if err != nil {
		return err
	}
	utils.Logger.Info().Msgf("Start method UpdateById %v", string(userJSON))

	if err := configs.Validator.Struct(updateUserDto); err != nil {
		return err
	}

	if updateUserDto.Password != nil && *updateUserDto.Password != "" {
		hashedPassword, err := authentication.HashPassword(*updateUserDto.Password)
		if err != nil {
			return err
		}
		updateUserDto.Password = &hashedPassword
	}

	return s.repo.UpdateById(updateUserDto, id)
}

func (s *service) Delete(id uint) error {
	utils.Logger.Info().Msgf("Start method Delete %d", id)
	return s.repo.Delete(id)
}
