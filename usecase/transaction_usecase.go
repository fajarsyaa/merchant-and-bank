package usecase

import (
	"project/model"
	"project/model/request"
	"project/repository"
	"project/utils"
	"time"
)

type TransactionUseCase interface {
	GetAllTransaction() []*model.TransactionModel
	InsertTransaction(tx *request.TransactionRequestModel) error
}

type transactionUsecaseImpl struct {
	transactionRepo repository.TransactionRepo
}

func (txUsecase *transactionUsecaseImpl) GetAllTransaction() []*model.TransactionModel {
	return txUsecase.transactionRepo.GetAllTransaction()
}

func (txUsecase *transactionUsecaseImpl) InsertTransaction(tx *request.TransactionRequestModel) error {
	transaction := &model.TransactionModel{}
	transaction.Id = utils.UuidGenerate()
	transaction.CreatedAt = time.Now()

	err := txUsecase.transactionRepo.InsertTransaction(transaction)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Internal Server Error",
		}
	}

	return nil
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepo) TransactionUseCase {
	return &transactionUsecaseImpl{
		transactionRepo: transactionRepo,
	}
}
