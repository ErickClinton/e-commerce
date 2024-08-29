package services

import (
	"eccomerce/configs"
	"eccomerce/internal/v1/cart"
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/user/dto"
	"eccomerce/internal/v1/user/repository"
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
	repo        repository.UserRepository
	cartService cart.CartService
}

func NewService(repo repository.UserRepository, cartService cart.CartService) Service {
	return &service{repo: repo, cartService: cartService}
}

func (s *service) Create(user *dto.CreateUserRequest) error {
	userJSON, _ := json.MarshalIndent(user, "", "    ")
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
		return errors.New("user already exist")
	}
	_, err = s.cartService.Create(entityUser.ID)
	if err != nil {
		return errors.New("error create wallet")
	}

	println()
	return err
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

	if err := configs.Validator.Struct(updateUserDto); err != nil {
		return err
	}

	if updateUserDto.Password != nil && *updateUserDto.Password != "" {
		hashedPassword, _ := authentication.HashPassword(*updateUserDto.Password)
		updateUserDto.Password = &hashedPassword
	}

	return s.repo.UpdateById(updateUserDto, id)
}

func (s *service) Delete(id uint) error {
	idJSON, _ := json.MarshalIndent(id, "", "    ")
	utils.Logger.Info().Msgf("Start method Delete %v", string(idJSON))
	return s.repo.Delete(id)
}
