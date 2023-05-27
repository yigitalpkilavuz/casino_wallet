package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Username     string `gorm:"size:255;index:idx_name,unique"`
	Password     string
	Balance      decimal.Decimal
	Transactions []Transaction `gorm:"foreignKey:WalletID"`
}
