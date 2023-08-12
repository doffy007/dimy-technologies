package repository

import (
	"context"
	"fmt"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/database/query"
	"github.com/doffy007/dimy-technologies/model"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	CreateTransaction(model.Transaction) error
}

type transactionRepository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

func (r repository) TransactionRepository() TransactionRepository {
	return &transactionRepository{
		ctx:    r.ctx,
		config: r.config,
		db:     r.db,
	}
}

// CreateTransaction implements TransactionRepository.
func (r transactionRepository) CreateTransaction(payload model.Transaction) error {
	_, err := r.db.NamedExecContext(r.ctx, query.CreateTransaction, payload)
	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository Create Transaction: %v\n", err.Error())
		return err
	}

	return nil
}
