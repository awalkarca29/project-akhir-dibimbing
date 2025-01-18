package entity

import "time"

type Product struct {
	ID            int
	Name          string
	Description   string
	Slug          string
	Price         int
	Stock         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Transaction   []Transaction
	ProductImages []ProductImage
}

type ProductImage struct {
	ID        int
	ProductID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
