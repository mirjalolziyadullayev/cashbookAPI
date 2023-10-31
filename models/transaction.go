package models

import "time"

type Transaction struct {
	ID        int
	UserID    int
	AccountID int
	Name      string
	Value     int64
	Type      bool
	Edited    time.Time
	Done      time.Time
}
