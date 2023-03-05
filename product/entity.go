package product

import "time"

type Product struct {
	ID          int
	Name        string
	Slug        string
	Description string
	Amount      int
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
