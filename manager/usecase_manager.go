package manager

import (
	"project/usecase"
	"sync"
)

type UsecaseManager interface {
	GetLoginUsecase() usecase.LoginUseCase
	GetCustomerUsecase() usecase.CustomerUseCase
	GetTransactionUsecase() usecase.TransactionUseCase
	GetMerchantUsecase() usecase.MerchantUseCase
}

type usecaseManager struct {
	repoManager     RepoManager
	loginUsecase    usecase.LoginUseCase
	customerUsecase usecase.CustomerUseCase
	txUsecase       usecase.TransactionUseCase
	merchUsecase    usecase.MerchantUseCase
}

var onceLoadLoginUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadTransactionUsecase sync.Once
var onceLoadMerchantUsecase sync.Once

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.loginUsecase = usecase.NewLoginUseCase(um.repoManager.GetLoginRepo())
	})
	return um.loginUsecase
}

func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUseCase {
	onceLoadCustomerUsecase.Do(func() {
		um.customerUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo())
	})
	return um.customerUsecase
}

func (um *usecaseManager) GetTransactionUsecase() usecase.TransactionUseCase {
	onceLoadTransactionUsecase.Do(func() {
		um.txUsecase = usecase.NewTransactionUseCase(um.repoManager.GetTransactionRepo(), um.repoManager.GetMerchantRepo(), um.repoManager.GetCustomerRepo())
	})
	return um.txUsecase
}

func (um *usecaseManager) GetMerchantUsecase() usecase.MerchantUseCase {
	onceLoadMerchantUsecase.Do(func() {
		um.merchUsecase = usecase.NewMerchantUseCase(um.repoManager.GetMerchantRepo())
	})
	return um.merchUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
