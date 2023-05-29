package storage

import (
	"fmt"

	"github.com/shopspring/decimal"
	entity "github.com/yigitalpkilavuz/casino_wallet/database/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Wallet{}, &entity.Transaction{})
	if err != nil {
		return err
	}
	return nil
}

func SeedData(db *gorm.DB) error {
	wallets := []entity.Wallet{
		{Username: "User1", Password: "password1", Balance: decimal.NewFromFloat(1000.0)},
		{Username: "User2", Password: "password2", Balance: decimal.NewFromFloat(2000.0)},
		{Username: "User3", Password: "password3", Balance: decimal.NewFromFloat(3000.0)},
		{Username: "User4", Password: "password4", Balance: decimal.NewFromFloat(4000.0)},
		{Username: "User5", Password: "password5", Balance: decimal.NewFromFloat(5000.0)},
	}
	for _, wallet := range wallets {
		result := db.FirstOrCreate(&wallet, wallet)
		if result.Error != nil {
			fmt.Println("Error inserting wallet:", result.Error)
			continue
		}
		if result.RowsAffected == 1 {
			fmt.Println("Inserted wallet:", wallet.Username)
		} else {
			fmt.Println("Wallet already exists:", wallet.Username)
		}

		transactions := []entity.Transaction{
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(100.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(200.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(300.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(400.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(500.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(600.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(700.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(800.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(900.0), Type: "debit"},
			{WalletID: wallet.ID, Amount: decimal.NewFromFloat(1000.0), Type: "debit"},
		}

		for _, transaction := range transactions {
			result := db.FirstOrCreate(&transaction, transaction)
			if result.Error != nil {
				fmt.Println("Error inserting transaction:", result.Error)
				continue
			}
			if result.RowsAffected == 1 {
				fmt.Println("Inserted transaction:", transaction.ID)
			} else {
				fmt.Println("Transaction already exists:", transaction.ID)
			}
		}
	}
	return nil
}
