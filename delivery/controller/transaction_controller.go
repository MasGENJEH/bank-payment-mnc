package controller

import (
	"fmt"
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

func NewTransactionsController(transactionUC usecase.TransactionsUsecase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *TransactionsController {
	return &TransactionsController{
		transactionUC: transactionUC,
		rg:         rg,
		authMiddleware: authMiddleware,
	}
}

func (t *TransactionsController) createHandler(ctx *gin.Context) {
	userId, exists := ctx.Get("user")
    if !exists {
        common.SendErrorResponse(ctx, http.StatusUnauthorized, "User not authorized")
        return
    }
	var payload entity.Transaction
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	customerID := fmt.Sprintf("%v", payload.CustomerId)
    if customerID != userId {
        common.SendErrorResponse(ctx, http.StatusForbidden, "Forbidden: Unauthorized customer_id")
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