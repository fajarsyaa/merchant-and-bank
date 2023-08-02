package usecase

import (
	"fmt"
	"project/model"
	"project/model/request"
	"project/repository"
	"project/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TransactionUseCase interface {
	GetAllTransaction() []*model.TransactionModel
	InsertTransaction(tx *request.TransactionRequestModel, ctx *gin.Context) error
}

type transactionUsecaseImpl struct {
	transactionRepo repository.TransactionRepo
	merchRepo       repository.MerchantRepo
	custRepo        repository.CustomerRepo
}

func (txUsecase *transactionUsecaseImpl) GetAllTransaction() []*model.TransactionModel {
	return txUsecase.transactionRepo.GetAllTransaction()
}

func (txUsecase *transactionUsecaseImpl) InsertTransaction(tx *request.TransactionRequestModel, ctx *gin.Context) error {
	transaction := &model.TransactionModel{}
	session := sessions.Default(ctx)
	existSession := session.Get("CustomerID")
	custId := existSession.(string)

	cust, err := txUsecase.custRepo.GetCustomerById(custId)
	if err != nil {
		fmt.Printf("%v", err)
		return &utils.AppError{
			ErrorCode:    400,
			ErrorMessage: "Unregistered Customer",
		}
	}

	if cust.Balance < tx.Amount {
		return &utils.AppError{
			ErrorCode:    403,
			ErrorMessage: "insufficient balance",
		}
	}

	merch, err := txUsecase.merchRepo.GetMerchantByNoRek(tx.MerchantNoRek)
	if err != nil {
		fmt.Printf("%v", err)
		return &utils.AppError{
			ErrorCode:    400,
			ErrorMessage: "Unregistered NoRek Merchant",
		}
	}

	err = txUsecase.custRepo.UpdateCustomerBalance(custId, tx.Amount)
	if err != nil {
		fmt.Printf("%v", err)
		return &utils.AppError{
			ErrorCode:    500,
			ErrorMessage: "Internal Server Error",
		}
	}

	err = txUsecase.merchRepo.UpdateMerchantBalance(merch.ID, tx.Amount)
	if err != nil {
		txUsecase.custRepo.Undo(custId, tx.Amount)
		fmt.Printf("%v", err)
		return &utils.AppError{
			ErrorCode:    500,
			ErrorMessage: "Internal Server Error",
		}
	}

	transaction.Id = utils.UuidGenerate()
	transaction.CreatedAt = time.Now()
	transaction.MerchantID = merch.ID
	transaction.CustomerID = custId
	transaction.Amount = tx.Amount

	err = txUsecase.transactionRepo.InsertTransaction(transaction)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Internal Server Error",
		}
	}

	return nil
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepo, merchRepo repository.MerchantRepo, custRepo repository.CustomerRepo) TransactionUseCase {
	return &transactionUsecaseImpl{
		transactionRepo: transactionRepo,
		merchRepo:       merchRepo,
		custRepo:        custRepo,
	}
}
