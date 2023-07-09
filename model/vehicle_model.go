package model

import "time"

type Vehicle struct {
	ID      uint
	Name    string
	Type    string
	Release time.Time
	Price   float64
	Status  string
}
