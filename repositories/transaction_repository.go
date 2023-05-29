package repository

type ITransactionRepository interface {
	IBaseRepository
}

type TransactionRepository struct {
	BaseRepository
}

func NewTransactionRepository(baseRepo BaseRepository) TransactionRepository {
	return TransactionRepository{BaseRepository: baseRepo}
}
