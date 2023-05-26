package service

import entity "github.com/yigitalpkilavuz/casino_wallet/entities"

type TransactionService interface {
	Find(id int) entity.Transaction
}

type transactionService struct {
}

func NewTransactionService() TransactionService {
	return &transactionService{}
}

func (service *transactionService) Find(id int) entity.Transaction {
	return entity.Transaction{}
}
