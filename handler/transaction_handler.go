package handler

import (
	"errors"
	"project/middleware"
	"project/model/request"
	"project/usecase"
	"project/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	txUsecase usecase.TransactionUseCase
}

func (txHandler TransactionHandler) GetAllTransaction(ctx *gin.Context) {
	transactions := txHandler.txUsecase.GetAllTransaction()
	if transactions == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "No Transaction Data",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"data":    transactions,
	})
}

func (txHandler TransactionHandler) CreateTransaction(ctx *gin.Context) {
	customer := &request.TransactionRequestModel{}
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if customer.MerchantNoRek == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "No rek Merchant cannot be empty",
		})
		return
	}
	if customer.Amount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Amount cannot be empty",
		})
		return
	}

	err = txHandler.txUsecase.InsertTransaction(customer, ctx)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred during create transaction",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"success": "success create transaction",
	})
}

func NewTransactionHandler(srv *gin.Engine, txUsecase usecase.TransactionUseCase) *TransactionHandler {
	txHandler := &TransactionHandler{
		txUsecase: txUsecase,
	}

	srv.POST("/transaction/create", middleware.RequireToken(), txHandler.CreateTransaction)
	srv.GET("/transactions", middleware.RequireToken(), txHandler.GetAllTransaction)

	return txHandler
}
