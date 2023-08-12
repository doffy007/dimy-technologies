package handler

type Handler interface {
	TransactionHandler() TransactionHandler
}
