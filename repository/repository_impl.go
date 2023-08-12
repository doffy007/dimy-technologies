package repository

type Repository interface {
	TransactionRepository() TransactionRepository
}
