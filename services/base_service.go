package service

import (
	"github.com/yigitalpkilavuz/casino_wallet/models"
	repository "github.com/yigitalpkilavuz/casino_wallet/repositories"
)

type BaseService struct {
	walletRepository   repository.WalletRepository
	transactionService repository.TransactionRepository
}

func NewBaseService(wallet repository.WalletRepository, transaction repository.TransactionRepository) BaseService {
	return BaseService{
		walletRepository:   wallet,
		transactionService: transaction,
	}

}

func ErrorResponse(status int, message string, description string) models.ErrorResponse {
	return models.ErrorResponse{
		Status:      status,
		Message:     message,
		Description: description,
	}
}
