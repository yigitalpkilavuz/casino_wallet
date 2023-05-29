package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	auth "github.com/yigitalpkilavuz/casino_wallet/auth"
	"github.com/yigitalpkilavuz/casino_wallet/caching"
	entity "github.com/yigitalpkilavuz/casino_wallet/database/entities"
	"github.com/yigitalpkilavuz/casino_wallet/models"
	model "github.com/yigitalpkilavuz/casino_wallet/models"
)

type IWalletService interface {
	Authenticate(req models.AuthenticateRequest) (models.AuthenticateResponse, models.ErrorResponse)
	Balance(walletID string) (models.BalanceResponse, models.ErrorResponse)
	Credit(req models.TransactionRequest) (models.TransactionResponse, models.ErrorResponse)
	Debit(req models.TransactionRequest) (models.TransactionResponse, models.ErrorResponse)
}

type WalletService struct {
	BaseService
	RedisService caching.RedisService
	Logger       *logrus.Logger
}

func NewWalletService(baseService BaseService, redisService caching.RedisService, logger *logrus.Logger) WalletService {
	return WalletService{
		BaseService:  baseService,
		RedisService: redisService,
		Logger:       logger,
	}
}

func (service *WalletService) Authenticate(req model.AuthenticateRequest) (model.AuthenticateResponse, model.ErrorResponse) {
	wallet, err := service.BaseService.walletRepository.GetWalletByUsername(req.Username)
	if err != nil {
		return model.AuthenticateResponse{}, ErrorResponse(422, "Something went wrong while finding wallet", err.Error())
	}

	if wallet.Password == req.Password {
		token, err := auth.CreateToken(req.Username)
		if err != nil {
			return model.AuthenticateResponse{}, ErrorResponse(500, "Could not generate token", err.Error())
		}
		return model.AuthenticateResponse{
			Username: wallet.Username,
			Balance:  wallet.Balance,
			Token:    token,
		}, model.ErrorResponse{}
	}
	walletJSON, _ := json.Marshal(wallet)
	service.RedisService.Set(fmt.Sprint(wallet.ID), string(walletJSON), time.Hour*1)
	return model.AuthenticateResponse{}, ErrorResponse(401, "Invalid Credentials", err.Error())
}

func (service *WalletService) Balance(id string) (model.BalanceResponse, model.ErrorResponse) {
	wallet := entity.Wallet{}
	err := service.BaseService.walletRepository.Get(id, &wallet)
	if err != nil {
		return model.BalanceResponse{}, ErrorResponse(422, "Something went wrong while finding wallet", err.Error())
	}
	return model.BalanceResponse{
		Username: wallet.Username,
		Balance:  wallet.Balance,
	}, model.ErrorResponse{}
}

func (service *WalletService) Credit(req model.TransactionRequest) (model.TransactionResponse, model.ErrorResponse) {
	transactionType := "credit"
	wallet := entity.Wallet{}
	err := service.BaseService.walletRepository.Get(req.WalletId, &wallet)
	if err != nil {
		return model.TransactionResponse{}, ErrorResponse(422, "Something went wrong while finding wallet", err.Error())
	}

	err = service.changeBalance(&wallet, req.Amount, transactionType)
	if err != nil {
		return model.TransactionResponse{}, ErrorResponse(422, "Something went wrong while updating balance", err.Error())
	}

	transaction := entity.Transaction{
		WalletID: wallet.ID,
		Amount:   req.Amount,
		Type:     transactionType,
	}

	err = service.BaseService.walletRepository.Create(&transaction)
	if err != nil {
		return model.TransactionResponse{}, ErrorResponse(422, "Something went wrong while updating balance", err.Error())
	}

	return model.TransactionResponse{
		Username:      wallet.Username,
		Balance:       wallet.Balance,
		TransactionID: transaction.ID,
	}, model.ErrorResponse{}
}

func (service *WalletService) Debit(req model.TransactionRequest) (model.TransactionResponse, model.ErrorResponse) {
	transactionType := "debit"
	wallet := entity.Wallet{}
	err := service.BaseService.walletRepository.Get(req.WalletId, &wallet)
	if err != nil {
		return model.TransactionResponse{}, ErrorResponse(422, "Something went wrong while finding wallet", err.Error())
	}

	if wallet.Balance.LessThan(req.Amount) {
		return model.TransactionResponse{}, ErrorResponse(403, "Insufficent funds", "")
	}
	err = service.changeBalance(&wallet, req.Amount, transactionType)
	if err != nil {
		return model.TransactionResponse{}, ErrorResponse(422, "Something went wrong while updating balance", err.Error())
	}

	transaction := entity.Transaction{
		WalletID: wallet.ID,
		Amount:   req.Amount,
		Type:     transactionType,
	}

	err = service.BaseService.walletRepository.Create(&transaction)
	if err != nil {
		return model.TransactionResponse{}, ErrorResponse(422, "Something went wrong while updating balance", err.Error())
	}

	return model.TransactionResponse{
		Username:      wallet.Username,
		Balance:       wallet.Balance,
		TransactionID: transaction.ID,
	}, model.ErrorResponse{}
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
	err := service.RedisService.Clear(fmt.Sprint(wallet.ID))
	if err != nil {
		service.Logger.Error("Could not clear cache for wallet %s: %v", wallet.Username, err)
	}

	walletJSON, _ := json.Marshal(wallet)
	service.RedisService.Set(fmt.Sprint(wallet.ID), string(walletJSON), time.Hour*1)
	return nil
}
