package entity

import "time"

type User struct {
	ID        int
	RoleID    int
	Role      Role
	Name      string
	Email     string
	Password  string
	Photo     string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
