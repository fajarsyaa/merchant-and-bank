package model

import "time"

type Credit struct {
	ID         uint
	VehicleID  uint
	CustomerID uint
	Price      float64
	Interest   float64
	DateIn     time.Time
	DateOut    time.Time
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
