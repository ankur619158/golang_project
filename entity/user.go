package entity

import "time"

type Person struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	PhoneNumber string
	City        string
	State       string
	Street1     string
	Street2     string
	ZipCode     string
}
