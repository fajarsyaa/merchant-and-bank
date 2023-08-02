package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"project/model"
)

type MerchantRepo interface {
	GetAllMerchant() []*model.MerchantModel
	GetMerchantByNoRek(NoRek string) (*model.MerchantModel, error)
	UpdateMerchantBalance(id string, balance int) error
}

type merchantRepoImpl struct {
	merchants []model.MerchantModel
}

func (merchRepo *merchantRepoImpl) GetAllMerchant() []*model.MerchantModel {
	result := make([]*model.MerchantModel, len(merchRepo.merchants))

	for i := range merchRepo.merchants {
		result[i] = &merchRepo.merchants[i]
	}

	return result
}

func (merchRepo *merchantRepoImpl) GetMerchantByNoRek(NoRek string) (*model.MerchantModel, error) {
	for _, merch := range merchRepo.merchants {
		if merch.NoRek == NoRek {
			return &merch, nil
		}
	}
	return nil, errors.New("not found")
}

func (merchRepo *merchantRepoImpl) UpdateMerchantBalance(id string, balance int) error {
	file, err := os.OpenFile("database/merchant.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed open file")
	}
	defer file.Close()

	for i, merch := range merchRepo.merchants {
		if merch.ID == id {
			merchRepo.merchants[i].Balance += balance
			break
		}
	}

	err = json.NewEncoder(file).Encode(merchRepo.merchants)
	if err != nil {
		fmt.Printf("%v", err)
		for i, merch := range merchRepo.merchants {
			if merch.ID == id {
				merchRepo.merchants[i].Balance -= balance
				break
			}
		}
		return errors.New("failed update json file")
	}

	err = file.Sync()
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed to sync JSON file")
	}

	return nil
}

func NewMerchantRepo() (MerchantRepo, error) {
	repo := &merchantRepoImpl{}

	file, err := os.Open("database/merchant.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&repo.merchants)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
