package model

import "time"

type Cash struct {
	ID          uint
	VehicleID   uint
	CustomerID  uint
	Price       float64
	DatePayment time.Time
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
