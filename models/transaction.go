package models

import "time"

type Transactions struct {
	ID              int
	UserID          int
	AccountID       int
	Name            string
	Value           float64
	TransactionType string
	Edited          time.Time
	Done            time.Time
}
