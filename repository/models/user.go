package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Person struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   null.Time
	Name        string
	PhoneNumber string
	City        string
	State       string
	Street1     string
	Street2     string
	ZipCode     string
}
