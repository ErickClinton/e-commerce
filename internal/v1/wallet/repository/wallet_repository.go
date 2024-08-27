package repository

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/pkg/utils"
	"gorm.io/gorm"
)

type WalletRepository interface {
	utils.Repository[entity.Wallet]
	UpdateBalance(userId uint, balance float64) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) Create(wallet *entity.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *walletRepository) GetByID(id uint) (*entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.First(&wallet, id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Adicione a implementação para Update
func (r *walletRepository) Update(entity *entity.Wallet) error {
	return r.db.Save(entity).Error
}

func (r *walletRepository) UpdateBalance(userId uint, balance float64) error {
	return r.db.Model(&entity.Wallet{}).Where("user_id = ?", userId).Update("balance", balance).Error
}

func (r *walletRepository) Delete(id uint) error {
	return r.db.Where("user_id = ?", id).Delete(&entity.Wallet{}).Error
}
