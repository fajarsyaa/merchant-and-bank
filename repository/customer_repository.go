package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"project/model"
	"project/utils"
)

type CustomerRepo interface {
	GetCustomerByUsername(username string) (*model.CustomerModel, error)
	InsertCustomer(customer *model.CustomerModel) error
	GetCustomerById(id string) (*model.CustomerModel, error)
	UpdateCustomerBalance(id string, balance int) error
	Undo(id string, balance int)
	TopUpBalance(id string, balance int) error
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
	return nil, errors.New("customer not found")
}

func (custRepo *customerRepoImpl) GetCustomerById(id string) (*model.CustomerModel, error) {
	for _, cust := range custRepo.customers {
		if cust.Id == id {
			return &cust, nil
		}
	}
	return nil, errors.New("customer not found")
}

func (custRepo *customerRepoImpl) TopUpBalance(id string, balance int) error {
	file, err := os.OpenFile("database/customer.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed open file")
	}
	defer file.Close()

	for i, cust := range custRepo.customers {
		if cust.Id == id {
			custRepo.customers[i].Balance += balance
			break
		}
	}

	err = json.NewEncoder(file).Encode(custRepo.customers)
	if err != nil {
		fmt.Printf("%v", err)
		for i, merch := range custRepo.customers {
			if merch.Id == id {
				custRepo.customers[i].Balance -= balance
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

func (custRepo *customerRepoImpl) UpdateCustomerBalance(id string, balance int) error {
	file, err := os.OpenFile("database/customer.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed open file")
	}
	defer file.Close()

	for i, cust := range custRepo.customers {
		if cust.Id == id {
			custRepo.customers[i].Balance -= balance
			break
		}
	}

	err = json.NewEncoder(file).Encode(custRepo.customers)
	if err != nil {
		fmt.Printf("%v", err)
		for i, merch := range custRepo.customers {
			if merch.Id == id {
				custRepo.customers[i].Balance += balance
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

func (custRepo *customerRepoImpl) Undo(id string, balance int) {
	file, _ := os.OpenFile("database/customer.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()

	for i, cust := range custRepo.customers {
		if cust.Id == id {
			custRepo.customers[i].Balance += balance
			break
		}
	}

	err := json.NewEncoder(file).Encode(custRepo.customers)
	if err != nil {
		fmt.Printf("%v", err)
		for i, merch := range custRepo.customers {
			if merch.Id == id {
				custRepo.customers[i].Balance -= balance
				break
			}
		}
	}

	err = file.Sync()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func (custRepo *customerRepoImpl) InsertCustomer(customer *model.CustomerModel) error {
	foundId := false
	for _, u := range custRepo.customers {
		if u.Id == customer.Id {
			foundId = true
			return errors.New("failed register")
		}
	}
	foundNoRek := false
	for _, u := range custRepo.customers {
		if u.NoRek == customer.NoRek {
			foundNoRek = true
			return &utils.AppError{
				ErrorCode:    500,
				ErrorMessage: "no rek is already used",
			}
		}
	}
	if !foundId && !foundNoRek {
		custRepo.customers = append(custRepo.customers, *customer)
	}

	file, err := os.OpenFile("database/customer.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed open file")
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(custRepo.customers)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed update json file")
	}

	err = file.Sync()
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed to sync JSON file")
	}

	return nil
}

func NewCustomerRepo() (CustomerRepo, error) {
	repo := &customerRepoImpl{}

	file, err := os.Open("database/customer.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&repo.customers)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
