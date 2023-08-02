package repository

import (
	"encoding/json"
	"errors"
	"os"
	"project/model"
)

type LoginRepo interface {
	GetCustomerByUsername(username string) (*model.CustomerModel, error)
}

type loginRepoImpl struct {
	customers []model.CustomerModel
}

func (lgnRepo *loginRepoImpl) GetCustomerByUsername(username string) (*model.CustomerModel, error) {
	for _, cust := range lgnRepo.customers {
		if cust.Username == username {
			return &cust, nil
		}
	}

	return nil, errors.New("customer not found")
}

func NewLoginRepo() (LoginRepo, error) {
	repo := &loginRepoImpl{}

	// Open the JSON file
	file, err := os.Open("database/customer.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.customers)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
