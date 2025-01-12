package entity

import "time"

type Order struct {
	ID            int
	ProductID     int
	UserID        int
	Quantity      int
	Total         int
	PaymentMethod string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
