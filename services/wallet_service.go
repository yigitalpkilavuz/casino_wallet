package service

import (
	"fmt"

	"github.com/shopspring/decimal"
	auth "github.com/yigitalpkilavuz/casino_wallet/auth"
	entity "github.com/yigitalpkilavuz/casino_wallet/database/entities"
	model "github.com/yigitalpkilavuz/casino_wallet/models"
	models "github.com/yigitalpkilavuz/casino_wallet/models"
)

type WalletService struct {
	BaseService
}

func NewWalletService(baseService BaseService) WalletService {
	return WalletService{
		BaseService: baseService,
	}
}

func (service *WalletService) Authenticate(req models.AuthenticateRequest) (models.AuthenticateResponse, models.ErrorResponse) {
	wallet, err := service.BaseService.walletRepository.GetWalletByUsername(req.Username)
	if err != nil {
		return models.AuthenticateResponse{}, ErrorResponse(500, "Something went wrong while finding wallet", err.Error())
	}

	if wallet.Password == req.Password {
		token, err := auth.CreateToken(req.Username)
		if err != nil {

			return models.AuthenticateResponse{}, ErrorResponse(500, "Could not generate token", err.Error())
		}
		return models.AuthenticateResponse{
			Username: wallet.Username,
			Balance:  wallet.Balance,
			Token:    token,
		}, models.ErrorResponse{}
	}
	return models.AuthenticateResponse{}, ErrorResponse(401, "Invalid Credentials", err.Error())
}

func (service *WalletService) Balance(id string) (models.BalanceResponse, models.ErrorResponse) {
	wallet := entity.Wallet{}
	err := service.BaseService.walletRepository.Get(id, &wallet)
	if err != nil {
		return models.BalanceResponse{}, ErrorResponse(500, "Something went wrong while finding wallet", err.Error())
	}
	return models.BalanceResponse{
		Username: wallet.Username,
		Balance:  wallet.Balance,
	}, models.ErrorResponse{}
}

func (service *WalletService) Credit(req model.TransactionRequest) (models.TransactionResponse, model.ErrorResponse) {
	transactionType := "credit"
	wallet := entity.Wallet{}
	err := service.BaseService.walletRepository.Get(req.WalletId, &wallet)
	if err != nil {
		return models.TransactionResponse{}, ErrorResponse(500, "Something went wrong while finding wallet", err.Error())
	}

	err = service.changeBalance(&wallet, req.Amount, transactionType)
	if err != nil {
		return models.TransactionResponse{}, ErrorResponse(500, "Something went wrong while updating balance", err.Error())
	}

	transaction := entity.Transaction{
		WalletID: wallet.ID,
		Amount:   req.Amount,
		Type:     transactionType,
	}

	err = service.BaseService.walletRepository.Create(&transaction)
	if err != nil {
		return models.TransactionResponse{}, ErrorResponse(500, "Something went wrong while updating balance", err.Error())
	}

	return models.TransactionResponse{
		Username:      wallet.Username,
		Balance:       wallet.Balance,
		TransactionID: transaction.ID,
	}, models.ErrorResponse{}
}

func (service *WalletService) Debit(req model.TransactionRequest) (model.TransactionResponse, model.ErrorResponse) {
	transactionType := "debit"
	wallet := entity.Wallet{}
	err := service.BaseService.walletRepository.Get(req.WalletId, &wallet)
	if err != nil {
		return models.TransactionResponse{}, ErrorResponse(500, "Something went wrong while finding wallet", err.Error())
	}

	err = service.changeBalance(&wallet, req.Amount, transactionType)
	if err != nil {
		return models.TransactionResponse{}, ErrorResponse(500, "Something went wrong while updating balance", err.Error())
	}

	transaction := entity.Transaction{
		WalletID: wallet.ID,
		Amount:   req.Amount,
		Type:     transactionType,
	}

	err = service.BaseService.walletRepository.Create(&transaction)
	if err != nil {
		return models.TransactionResponse{}, ErrorResponse(500, "Something went wrong while updating balance", err.Error())
	}

	return models.TransactionResponse{
		Username:      wallet.Username,
		Balance:       wallet.Balance,
		TransactionID: transaction.ID,
	}, models.ErrorResponse{}
}

func (service *WalletService) changeBalance(wallet *entity.Wallet, amount decimal.Decimal, transaction string) error {
	switch transaction {
	case "debit":
		wallet.Balance = wallet.Balance.Sub(amount)
	case "credit":
		wallet.Balance = wallet.Balance.Add(amount)
	}
	if err := service.BaseService.walletRepository.Update(fmt.Sprint(wallet.ID), &wallet); err != nil {
		return err
	}
	return nil
}
