package usecase

import (
	"project/model"
	"project/repository"
)

type MerchantUseCase interface {
	GetAllMerchant() []*model.MerchantModel
}

type merchantUsecaseImpl struct {
	merchRepo repository.MerchantRepo
}

func (merchUsecase *merchantUsecaseImpl) GetAllMerchant() []*model.MerchantModel {
	return merchUsecase.merchRepo.GetAllMerchant()
}

func NewMerchantUseCase(merchRepo repository.MerchantRepo) MerchantUseCase {
	return &merchantUsecaseImpl{
		merchRepo: merchRepo,
	}
}
