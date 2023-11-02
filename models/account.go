package models

import "time"

type Account struct {
	ID      int
	UserID  int
	Name    string
	Created time.Time
	Edited  time.Time
	Transactions []struct {
		ID              int
		UserID          int
		AccountID       int
		Name            string
		Value           int64
		TransactionType bool
		Edited          time.Time
		Done            time.Time
	}
}

