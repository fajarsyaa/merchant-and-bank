package model

import "time"

type Rent struct {
	ID         uint
	VehicleID  uint
	CustomerID uint
	Price      float64
	DateIn     time.Time
	DateOut    time.Time
	Status     string
	CreatedBy  string
	UpdatedBy  string
}
