package models

import "time"

type Account struct {
	ID int
	UserID int
	Name string
	Since time.Time
	Edited time.Time
	Transactions []Transaction
}