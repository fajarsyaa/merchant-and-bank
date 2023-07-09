package model

import "time"

type Customer struct {
	ID        uint
	FullName  string
	Password  string
	NIK       string
	NoPhone   string
	Email     string
	Address   string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
