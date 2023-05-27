package entity

import (
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Entity
	WalletID uint
	Amount   decimal.Decimal
	Type     string
}
