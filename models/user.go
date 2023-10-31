package models

import "time"

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Edit      time.Time
	Since     time.Time
	Account   []Account
}
