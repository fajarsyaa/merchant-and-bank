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

type LoginHandler struct {
	lgUsecase usecase.LoginUseCase
}

func (lgHandler LoginHandler) Login(ctx *gin.Context) {
	loginReq := &request.LoginRequestModel{}
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if loginReq.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name cannot be empty",
		})
		return
	}
	if loginReq.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Password cannot be empty",
		})
		return
	}

	usr, err := lgHandler.lgUsecase.Login(loginReq, ctx)

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
				"errorMessage": "An error occurred during login",
			})
		}
		return
	}

	if usr == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Name is not registered",
		})
		return
	}

	tokenJwt, err := utils.GenerateToken(loginReq.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid Token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  tokenJwt,
	})
}

func (lgHandler LoginHandler) Logout(ctx *gin.Context) {
	err := lgHandler.lgUsecase.Logout(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Unauthorize",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Logout",
	})
}

func NewLoginHandler(srv *gin.Engine, lgUsecase usecase.LoginUseCase) *LoginHandler {
	lgHandler := &LoginHandler{
		lgUsecase: lgUsecase,
	}

	srv.POST("/login", lgHandler.Login)
	srv.POST("/logout", middleware.RequireToken(), lgHandler.Logout)

	return lgHandler
}
