package entity

import "time"

type User struct {
	ID        int
	RoleId    int
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
