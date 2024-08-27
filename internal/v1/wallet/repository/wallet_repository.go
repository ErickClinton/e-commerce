package repository

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/pkg/utils"
	"gorm.io/gorm"
)

type WalletRepository interface {
	utils.Repository[entity.Wallet]
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

func (r *walletRepository) Update(wallet *entity.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *walletRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Wallet{}, id).Error
}
