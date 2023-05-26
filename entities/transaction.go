package entity

type Transaction struct {
	Id       int
	WalletId int
	Amount   float32
	Type     string
	// CreatedAt date
}
