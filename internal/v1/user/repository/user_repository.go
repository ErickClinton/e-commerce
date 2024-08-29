package repository

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/user/dto"
	"eccomerce/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	utils.Repository[entity.User]
	GetByEmail(email string) (*entity.User, error)
	UpdateById(user *dto.UpdateUserRequest, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.Model(&entity.User{}).Preload("Products").Preload("Wallet").Preload("Cart").Where("id = ?", id).Find(&user).Error; err != nil {

		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) UpdateById(updateUserDto *dto.UpdateUserRequest, id int) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(updateUserDto).First(&updateUserDto, id).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
