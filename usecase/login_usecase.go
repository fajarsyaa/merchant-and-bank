package usecase

import (
	"fmt"
	"project/middleware"
	"project/model"
	"project/model/request"
	"project/repository"
	"project/utils"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Login(cust *request.LoginRequestModel, ctx *gin.Context) (*model.CustomerModel, error)
	Logout(ctx *gin.Context) error
}

type loginUsecase struct {
	loginRepo repository.LoginRepo
}

func (loginUsecase *loginUsecase) Login(cust *request.LoginRequestModel, ctx *gin.Context) (*model.CustomerModel, error) {
	// Login session
	session := sessions.Default(ctx)
	existSession := session.Get("Username")
	if existSession != nil {
		return nil, &utils.AppError{
			ErrorCode:    403,
			ErrorMessage: fmt.Sprintf("You are already logged in as %v", existSession),
		}
	}

	existData, err := loginUsecase.loginRepo.GetCustomerByUsername(cust.Username)
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    400,
			ErrorMessage: "Username is not registered",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(existData.Password), []byte(cust.Password))
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    400,
			ErrorMessage: "Password does not match",
		}
	}

	// Login session
	session.Set("Username", existData.Username)
	session.Set("CustomerID", existData.Id)
	session.Save()

	existData.Password = ""
	return existData, nil
}

func (loginUsecase *loginUsecase) Logout(ctx *gin.Context) error {
	h := &middleware.AuthHeader{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		return &utils.AppError{
			ErrorCode:    403,
			ErrorMessage: "Unautorized",
		}
	}
	tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
	loginUsecase.loginRepo.TokenBlock(tokenString)

	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	return nil
}

func NewLoginUseCase(loginRepo repository.LoginRepo) LoginUseCase {
	return &loginUsecase{
		loginRepo: loginRepo,
	}
}
