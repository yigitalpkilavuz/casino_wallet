package models

import "github.com/shopspring/decimal"

type AuthenticateResponse struct {
	Username string
	Balance  decimal.Decimal
	Token    string
}

type BalanceResponse struct {
	Username string
	Balance  decimal.Decimal
}

type TransactionResponse struct {
	Username      string
	Balance       decimal.Decimal
	TransactionID uint
}
