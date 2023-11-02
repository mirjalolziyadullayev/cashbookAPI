package models

import "time"

type Transactions struct {
	ID              int
	UserID          int
	AccountID       int
	Name            string
	Value           int64
	TransactionType bool
	Edited          time.Time
	Done            time.Time
}
