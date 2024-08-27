package services

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/wallet/dto"
	"eccomerce/internal/v1/wallet/repository"
	"eccomerce/pkg/utils"
)

type WalletService interface {
	utils.Service[dto.CreateWalletRequest, entity.Wallet]
}

type walletService struct {
	repo repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) WalletService {
	return &walletService{repo: repo}
}

func (s *walletService) Create(wallet *dto.CreateWalletRequest) error {
	utils.Logger.Info().Msgf("Start method create wallet for user %d", wallet.UserId)
	entityWallet := &entity.Wallet{
		UserId:  wallet.UserId,
		Balance: wallet.Balance, // Saldo inicial passado no request
	}
	return s.repo.Create(entityWallet)
}

func (s *walletService) GetByID(id uint) (*entity.Wallet, error) {
	return s.repo.GetByID(id)
}
func (s *walletService) Update(wallet *dto.CreateWalletRequest) error {

	return s.repo.UpdateBalance(wallet.UserId, wallet.Balance)

}
func (s *walletService) Delete(id uint) error {
	return s.repo.Delete(id)
}
