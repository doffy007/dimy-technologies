package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doffy007/dimy-technologies/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUsecase struct {
	mock.Mock
}

func (m *MockTransactionUsecase) RegisterTransaction(params model.Transaction) (bool, model.BaseResponse) {
	args := m.Called(params)
	return args.Bool(0), args.Get(1).(model.BaseResponse)
}

func TestTransactionRegister_Failure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	mockUsecase := new(MockTransactionUsecase)
	mockUsecase.On("RegisterTransaction", mock.Anything).Return(false, model.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "error message",
		Data:       nil,
	})

	handler := transactionHandler{
		usecase: mockUsecase,
	}

	router.POST("/transaction", handler.TransactionRegister)

	req, _ := http.NewRequest(http.MethodPost, "/transaction", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
