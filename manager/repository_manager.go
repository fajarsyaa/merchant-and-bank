package manager

type RepoManager interface {
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
