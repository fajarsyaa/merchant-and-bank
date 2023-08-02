package handler

import (
	"errors"
	"project/model"
	"project/usecase"
	"project/utils"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	CustUsecase usecase.CustomerUseCase
}

func (custHandler CustomerHandler) CreateCustomer(ctx *gin.Context) {
	customer := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if customer.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name cannot be empty",
		})
		return
	}
	if customer.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Password cannot be empty",
		})
		return
	}
	if customer.NoRek == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "NoRek cannot be empty",
		})
		return
	}
	if customer.NoRek == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Saldo cannot be empty",
		})
		return
	}

	err = custHandler.CustUsecase.InsertCustomer(customer)

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
				"errorMessage": "An error occurred during register",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"success": "success register",
	})
}

func (custHandler CustomerHandler) TopUpBalance(ctx *gin.Context) {
	customer := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	session := sessions.Default(ctx)
	existSession := session.Get("CustomerID")
	if existSession == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Unautorized",
		})
		return
	}

	if customer.Balance == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Balance cannot be empty",
		})
		return
	}

	customer.Id = existSession.(string)
	custHandler.CustUsecase.TopUpBalance(customer)
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
				"errorMessage": "An error occurred during TopUp",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"success": "success Top Up",
	})
}

func NewCustomerHandler(srv *gin.Engine, CustUsecase usecase.CustomerUseCase) *CustomerHandler {
	CustHandler := &CustomerHandler{
		CustUsecase: CustUsecase,
	}

	srv.POST("/register", CustHandler.CreateCustomer)
	srv.POST("/topup", CustHandler.TopUpBalance)

	return CustHandler
}
