package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	WalletID uint
	Amount   decimal.Decimal
	Type     string
}
