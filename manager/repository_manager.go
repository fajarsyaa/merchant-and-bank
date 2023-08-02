package manager

import (
	"project/repository"
	"sync"
)

var err error

type RepoManager interface {
	GetLoginRepo() repository.LoginRepo
	GetCustomerRepo() repository.CustomerRepo
	GetTransactionRepo() repository.TransactionRepo
}

type repoManager struct {
	loginRepo    repository.LoginRepo
	customerRepo repository.CustomerRepo
	txRepo       repository.TransactionRepo
}

var onceLoadLoginRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadTransactionRepo sync.Once

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
			panic(err)
		}
	})
	return rm.customerRepo
}

func (rm *repoManager) GetTransactionRepo() repository.TransactionRepo {
	onceLoadTransactionRepo.Do(func() {
		rm.txRepo, err = repository.NewTransactionRepo()
		if err != nil {
			panic(err)
		}
	})
	return rm.txRepo
}

func NewRepoManager() RepoManager {
	return &repoManager{}
}
