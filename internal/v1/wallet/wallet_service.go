package wallet

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/wallet/dto"
	"eccomerce/pkg/utils"
)

type WalletService interface {
	utils.Service[dto.CreateWalletRequest, entity.Wallet]
}

type walletService struct {
	repo WalletRepository
}

func NewWalletService(repo WalletRepository) WalletService {
	return &walletService{repo: repo}
}

func (s *walletService) Create(wallet *dto.CreateWalletRequest) error {
	utils.Logger.Info().Msgf("Start method create wallet for user %d", wallet.UserId)
	entityWallet := &entity.Wallet{
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
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
