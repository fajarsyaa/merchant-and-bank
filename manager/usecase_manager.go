package manager

type UsecaseManager interface {
}

type usecaseManager struct {
	repoManager RepoManager
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
