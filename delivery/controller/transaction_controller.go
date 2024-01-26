package controller

import (
	"net/http"
	"test-mnc/config"
	"test-mnc/delivery/middleware"
	"test-mnc/entity"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	transactionUC  usecase.TransactionsUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func NewTransactionsController(transactionUC usecase.TransactionsUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware,) *TransactionsController {
	return &TransactionsController{
		transactionUC: transactionUC,
		rg:         rg,
		authMiddleware: authMiddleware,
	}
}

func (t *TransactionsController) createHandler(ctx *gin.Context) {
	var payload entity.Transaction
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactions, err := t.transactionUC.RequestNewPayment(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Created")
}

func (t *TransactionsController) Route() {
	t.rg.POST(config.Payment, t.authMiddleware.RequireToken(), t.createHandler)
}