package entity

import (
	"github.com/shopspring/decimal"
)

type Wallet struct {
	Entity
	Username     string `gorm:"size:255;index:idx_name,unique"`
	Password     string
	Balance      decimal.Decimal
	Transactions []Transaction `gorm:"foreignKey:WalletID"`
}
