package usecase

import (
	"fmt"
	"project/model"
	"project/repository"
	"project/utils"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase interface {
	InsertCustomer(cust *model.CustomerModel) error
}

type customerUsecaseImpl struct {
	customerRepo repository.CustomerRepo
}

func (custUsecase *customerUsecaseImpl) InsertCustomer(cust *model.CustomerModel) error {
	cust.Id = utils.UuidGenerate()
	passHash, err := bcrypt.GenerateFromPassword([]byte(cust.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("customerUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	cust.Password = string(passHash)

	err = custUsecase.customerRepo.InsertCustomer(cust)
	if err != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Internal Server Error",
		}
	}

	return nil
}

func NewCustomerUseCase(customerRepo repository.CustomerRepo) CustomerUseCase {
	return &customerUsecaseImpl{
		customerRepo: customerRepo,
	}
}
