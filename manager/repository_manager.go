package manager

import (
	"project/repository"
	"sync"
)

type RepoManager interface {
	GetLoginRepo() repository.LoginRepo
}

type repoManager struct {
	infraManager InfraManager
	loginRepo    repository.LoginRepo
}

var onceLoadLoginRepo sync.Once

func (rm *repoManager) GetLoginRepo() repository.LoginRepo {
	onceLoadLoginRepo.Do(func() {
		rm.loginRepo = repository.NewLoginRepo(rm.infraManager.GetDB())
	})
	return rm.loginRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
