package entity

import "time"

type User struct {
	ID        int
	RoleID    int
	Roles     Role `gorm:"foreignKey:RoleID"`
	Name      string
	Email     string
	Password  string
	Photo     string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
