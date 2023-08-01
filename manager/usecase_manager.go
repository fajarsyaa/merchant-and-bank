package manager

import (
	"project/usecase"
	"sync"
)

type UsecaseManager interface {
	GetLoginUsecase() usecase.LoginUseCase
}

type usecaseManager struct {
	repoManager  RepoManager
	loginUsecase usecase.LoginUseCase
}

var onceLoadLoginUsecase sync.Once

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.loginUsecase = usecase.NewLoginUseCase(um.repoManager.GetLoginRepo())
	})
	return um.loginUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
