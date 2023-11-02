package models

import "time"

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Edit      time.Time
	Created   time.Time
	Account   []struct {
		ID           int
		UserID       int
		Name         string
		Created      time.Time
		Edited       time.Time
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
}
