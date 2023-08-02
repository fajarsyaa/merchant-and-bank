package repository

import (
	"encoding/json"
	"errors"
	"os"
	"project/model"
)

type CustomerRepo interface {
	GetCustomerByUsername(username string) (*model.CustomerModel, error)
	InsertCustomer(customer *model.CustomerModel) error
}

type customerRepoImpl struct {
	customers []model.CustomerModel
}

func (custRepo *customerRepoImpl) GetCustomerByUsername(username string) (*model.CustomerModel, error) {
	for _, cust := range custRepo.customers {
		if cust.Username == username {
			return &cust, nil
		}
	}
	return nil, errors.New("user not found")
}

func (custRepo *customerRepoImpl) InsertCustomer(customer *model.CustomerModel) error {
	found := false
	for i, u := range custRepo.customers {
		if u.Id == customer.Id {
			custRepo.customers[i] = *customer
			found = true
			break
		}
	}
	if !found {
		custRepo.customers = append(custRepo.customers, *customer)
	}

	file, err := os.OpenFile("database/customer.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(custRepo.customers)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepo() (CustomerRepo, error) {
	repo := &customerRepoImpl{}

	// Open the JSON file
	file, err := os.Open("database/customer.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the file into the customers slice
	err = json.NewDecoder(file).Decode(&repo.customers)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
