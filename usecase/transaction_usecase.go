package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/model"
	"github.com/doffy007/dimy-technologies/repository"
)

type TransactionUsecase interface {
	RegisterTransaction(model.Transaction) (bool, model.BaseResponse)
}

type transactionUsecase struct {
	ctx                   context.Context
	config                *config.Config
	transactionRepository repository.TransactionRepository
}

func (u usecase) TransactionHandler() TransactionUsecase {
	return &transactionUsecase{
		ctx:                   u.ctx,
		config:                u.config,
		transactionRepository: u.repository.TransactionRepository(),
	}
}

// RegisterTransaction implements TransactionUsecase.
func (u transactionUsecase) RegisterTransaction(params model.Transaction) (bool, model.BaseResponse) {
	var payload model.Transaction
	rec, _ := json.Marshal(params)
	json.Unmarshal(rec, &payload)

	err := u.transactionRepository.CreateTransaction(payload)

	if err != nil {
		return false, model.BaseResponse{}.InternalServerError(err.Error())
	}

	return true, model.BaseResponse{StatusCode: http.StatusCreated, Message: "success created user", Data: true}
}
