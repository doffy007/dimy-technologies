package usecase

type Usecase interface {
	TransactionHandler() TransactionUsecase
}
