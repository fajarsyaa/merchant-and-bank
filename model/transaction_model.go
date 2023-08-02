package model

import "time"

type TransactionModel struct {
	Id         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	MerchantID string    `json:"merchant_id"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}
