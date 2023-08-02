package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"project/model"
)

type TransactionRepo interface {
	GetAllTransaction() []*model.TransactionModel
	InsertTransaction(transaction *model.TransactionModel) error
}

type transactionRepoImpl struct {
	transactions []model.TransactionModel
}

func (custRepo *transactionRepoImpl) GetAllTransaction() []*model.TransactionModel {
	result := make([]*model.TransactionModel, len(custRepo.transactions))

	for i := range custRepo.transactions {
		result[i] = &custRepo.transactions[i]
	}

	return result
}

func (custRepo *transactionRepoImpl) InsertTransaction(transaction *model.TransactionModel) error {

	custRepo.transactions = append(custRepo.transactions, *transaction)

	file, err := os.OpenFile("database/transaction.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.New("failed open file")
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(custRepo.transactions)
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

func NewTransactionRepo() (TransactionRepo, error) {
	repo := &transactionRepoImpl{}

	file, err := os.Open("database/transaction.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&repo.transactions)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
