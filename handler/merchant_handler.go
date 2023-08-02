package handler

import (
	"net/http"
	"project/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	merchUsecase usecase.MerchantUseCase
}

func (merchHandler MerchantHandler) GetAllMerchant(ctx *gin.Context) {
	merchants := merchHandler.merchUsecase.GetAllMerchant()
	if merchants == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "No Merchant Data",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"data":    merchants,
	})
}

func NewMerchantHandler(srv *gin.Engine, merchUsecase usecase.MerchantUseCase) *MerchantHandler {
	merchHandler := &MerchantHandler{
		merchUsecase: merchUsecase,
	}

	srv.GET("/merchants", merchHandler.GetAllMerchant)

	return merchHandler
}
