package payment

import (
	"api-ecommerce/transaction"
	"time"
)

type Payment struct {
	ID            int
	TransactionID int
	UserID        int
	InvoiceCode   string
	Email         string
	Name          string
	TotalAmount   int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Transaction   transaction.Transaction `gorm:"references:TransactionID`
}
