package handler

import (
	"fmt"
	"net/http"

	"github.com/doffy007/dimy-technologies/helper"
	"github.com/doffy007/dimy-technologies/model"
	"github.com/doffy007/dimy-technologies/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	TransactionRegister(c *gin.Context)
}

type transactionHandler struct {
	usecase usecase.TransactionUsecase
}

func (h handler) TransactionHandler() TransactionHandler {
	return transactionHandler{
		usecase: h.usecase.TransactionHandler(),
	}
}

// TransactionRegister implements TransactionHandler.
func (h transactionHandler) TransactionRegister(c *gin.Context) {
	var params model.Transaction

	if err := c.ShouldBind(&params); err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler TransactionRegister Parse Param: %v\n", err.Error())
		helper.NewErrorResponse(c, http.StatusBadRequest, "Error parse params", err.Error())
		return
	}

	ok, res := h.usecase.RegisterTransaction(params)

	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler TransactionRegister From Usecase: %v\n", res.Message)
		helper.NewErrorResponse(c, res.StatusCode, res.Message, []string{res.Message})
		return
	}

	helper.NewSuccessResponse(c, res.StatusCode, res.Message, res.Data)
}
