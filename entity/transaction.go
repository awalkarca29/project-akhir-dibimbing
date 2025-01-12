package entity

import (
	"time"
)

type Transaction struct {
	ID            int
	ProductID     int
	UserID        int
	Quantity      int
	Total         int
	PaymentMethod string
	Status        string
	User          User
	Product       Product
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
