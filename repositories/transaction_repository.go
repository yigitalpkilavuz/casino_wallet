package repository

type ITransactionRepository interface {
}

type TransactionRepository struct {
	BaseRepository
}

func NewTransactionRepository(baseRepo BaseRepository) TransactionRepository {
	return TransactionRepository{BaseRepository: baseRepo}
}
