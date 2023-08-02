package manager

import (
	"errors"
	"project/repository"
	"sync"
)

var err error

type RepoManager interface {
	GetLoginRepo() repository.LoginRepo
	GetCustomerRepo() repository.CustomerRepo
	GetTransactionRepo() repository.TransactionRepo
	GetMerchantRepo() repository.MerchantRepo
}

type repoManager struct {
	loginRepo    repository.LoginRepo
	customerRepo repository.CustomerRepo
	txRepo       repository.TransactionRepo
	merchRepo    repository.MerchantRepo
}

var onceLoadLoginRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadTransactionRepo sync.Once
var onceLoadMerchantRepo sync.Once

func (rm *repoManager) GetLoginRepo() repository.LoginRepo {
	onceLoadLoginRepo.Do(func() {
		rm.loginRepo, err = repository.NewLoginRepo()
		if err != nil {
			panic(err)
		}
	})
	return rm.loginRepo
}

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.customerRepo, err = repository.NewCustomerRepo()
		if err != nil {
			if err != errors.New("EOF") {
				panic(err)
			}
		}
	})
	return rm.customerRepo
}

func (rm *repoManager) GetTransactionRepo() repository.TransactionRepo {
	onceLoadTransactionRepo.Do(func() {
		rm.txRepo, err = repository.NewTransactionRepo()
		if err != nil {
			if err != errors.New("EOF") {
				panic(err)
			}
		}
	})
	return rm.txRepo
}

func (rm *repoManager) GetMerchantRepo() repository.MerchantRepo {
	onceLoadMerchantRepo.Do(func() {
		rm.merchRepo, err = repository.NewMerchantRepo()
		if err != nil {
			if err != errors.New("EOF") {
				panic(err)
			}
		}
	})
	return rm.merchRepo
}

func NewRepoManager() RepoManager {
	return &repoManager{}
}
