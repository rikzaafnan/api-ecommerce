package transaction

import "time"

type Transaction struct {
	ID                 int
	Name               string
	Email              string
	Code               string
	Address            string
	TransactionStatus  string
	TransactionTotal   int
	Description        string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	TransactionDetails []TransactionDetail `gorm:"foreignKey:ID`
}

type TransactionDetail struct {
	ID            int
	TransactionID int
	ProductID     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
