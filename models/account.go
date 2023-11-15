package models

import "time"

type Account struct {
	ID           int
	UserID       int
	Name         string
	Balance      float64
	Created      time.Time
	Edited       time.Time
	Transactions []struct {
		ID              int
		UserID          int
		AccountID       int
		Name            string
		Value           float64
		TransactionType string
		Edited          time.Time
		Done            time.Time
	}
}
