package models

import "github.com/shopspring/decimal"

type AuthenticateRequest struct {
	Username string
	Password string
}

type TransactionRequest struct {
	Amount   decimal.Decimal
	WalletId string
}
