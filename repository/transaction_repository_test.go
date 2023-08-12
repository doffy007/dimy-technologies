package repository

import (
	"context"
	"testing"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/model"
	"github.com/stretchr/testify/assert"
)

func Test_repository_CreateTransaction(t *testing.T) {
	repo := NewRepository(context.Background(), &config.Config{})

	testPayload := model.Transaction{
		CustomerAddressID: 1,
		Product:           2,
		PaymentMethod:     2,
		ProductQty:        2,
	}

	err := repo.TransactionRepository().CreateTransaction(testPayload)

	assert.NoError(t, err)
}
