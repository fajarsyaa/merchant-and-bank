package usecase

import (
	"errors"
	"fmt"
	"project/model"
	"project/repository"
	"project/utils"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase interface {
	InsertCustomer(cust *model.CustomerModel) error
	TopUpBalance(cust *model.CustomerModel) error
	GetCustomerById(id string) (*model.CustomerModel, error)
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
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			return &utils.AppError{
				ErrorCode:    400,
				ErrorMessage: appError.ErrorMessage,
			}
		} else {
			return &utils.AppError{
				ErrorCode:    500,
				ErrorMessage: "Internal Server Error",
			}
		}
	}

	return nil
}

func (custUsecase *customerUsecaseImpl) GetCustomerById(id string) (*model.CustomerModel, error) {
	return custUsecase.customerRepo.GetCustomerById(id)
}

func (custUsecase *customerUsecaseImpl) TopUpBalance(cust *model.CustomerModel) error {
	err := custUsecase.customerRepo.TopUpBalance(cust.Id, cust.Balance)
	if err != nil {
		fmt.Printf("%v", err)
		return &utils.AppError{
			ErrorCode:    500,
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
