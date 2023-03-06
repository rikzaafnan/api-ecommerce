package payment

import (
	"api-ecommerce/transaction"

	"gorm.io/gorm"
)

type RepositoryPayment interface {
	Save(input Payment) (Payment, error)
	FindByID(paymentID int) (Payment, error)
	FindAll(userID int) ([]Payment, error)
}

type repositoryPaymentImpl struct {
	db *gorm.DB
}

func NewRepositoryPayment(db *gorm.DB) *repositoryPaymentImpl {
	return &repositoryPaymentImpl{db}
}

func (r *repositoryPaymentImpl) Save(input Payment) (Payment, error) {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return input, err
	}

	if err := tx.Create(&input).Error; err != nil {
		tx.Rollback()
		return input, err
	}

	if err := tx.Model(&transaction.Transaction{}).Where("id = ?", input.TransactionID).Update("transaction_status", "paid").Error; err != nil {
		tx.Rollback()
		return input, err
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return input, err
	}

	return input, nil

}

func (r *repositoryPaymentImpl) FindByID(paymentID int) (Payment, error) {

	var payment Payment

	err := r.db.Where("id = ?", paymentID).Preload("Transaction.TransactionDetails").First(&payment).Error
	if err != nil {

		return payment, err
	}

	return payment, nil

}

func (r *repositoryPaymentImpl) FindAll(userID int) ([]Payment, error) {

	var payments []Payment

	if userID != 0 {

		err := r.db.Model(&Payment{}).Where("user_id = ?", userID).Preload("Transaction.TransactionDetails").Find(&payments).Error
		if err != nil {

			return payments, err
		}

	} else {
		err := r.db.Model(&Payment{}).Preload("Transaction.TransactionDetails").Find(&payments).Error
		if err != nil {

			return payments, err
		}
	}

	return payments, nil

}
