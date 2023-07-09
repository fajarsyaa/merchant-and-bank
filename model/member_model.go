package model

import "time"

type Member struct {
	ID        uint
	UserID    uint
	Type      string
	Expire    time.Time
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
