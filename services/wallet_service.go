package service

import entity "github.com/yigitalpkilavuz/casino_wallet/entities"

type WalletService interface {
	Find(id int) (entity.Wallet, error)
}

type walletService struct {
}

func NewUserService() WalletService {
	return &walletService{}
}

func (service *walletService) Find(id int) (entity.Wallet, error) {
	return entity.Wallet{}, nil
}
