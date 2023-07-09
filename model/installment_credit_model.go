package model

import "time"

type InstallmentCredit struct {
	ID              uint
	VehicleID       uint
	CreditID        uint
	Price           float64
	TotalPaymentNow float64
	DatePayment     time.Time
	DateFinish      time.Time
	DueDate         int
	Status          bool
	Suspend         bool
}
