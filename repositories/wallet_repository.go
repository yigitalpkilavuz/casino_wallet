package repository

import entity "github.com/yigitalpkilavuz/casino_wallet/database/entities"

type WalletRepository struct {
	BaseRepository
}

func NewWalletRepository(baseRepo BaseRepository) WalletRepository {
	return WalletRepository{BaseRepository: baseRepo}
}

func (r *WalletRepository) GetWalletByUsername(username string) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.Where("username = ?", username).First(&wallet).Error; err != nil {
		return entity.Wallet{}, err
	}
	return wallet, nil
}
