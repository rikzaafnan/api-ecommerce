package transaction

import (
	"fmt"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	Save(input Transaction) (Transaction, error)
	FindByID(transactionID int) (Transaction, error)
	FindAll(userID int) ([]Transaction, error)
}

type repositoryTransactionImpl struct {
	db *gorm.DB
}

func NewRepositoryTransaction(db *gorm.DB) *repositoryTransactionImpl {
	return &repositoryTransactionImpl{db}
}

func (r *repositoryTransactionImpl) Save(transaction Transaction) (Transaction, error) {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return transaction, err
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return transaction, err
	}

	fmt.Println(transaction.TransactionDetails)
	fmt.Println(len(transaction.TransactionDetails))

	// for _, inputTransactionDetail := range transaction.TransactionDetails {

	// 	var transactionDetail TransactionDetail
	// 	transactionDetail.TransactionID = transaction.ID
	// 	transactionDetail.ProductID = inputTransactionDetail.ProductID

	// 	fmt.Println(inputTransactionDetail)

	// 	if err := tx.Create(&transactionDetail).Error; err != nil {
	// 		tx.Rollback()
	// 		return transaction, err
	// 	}

	// }

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	fmt.Println("success")
	fmt.Println(transaction)

	return transaction, nil

}

func (r *repositoryTransactionImpl) FindByID(transactionID int) (Transaction, error) {

	var transaction Transaction

	err := r.db.Where("id = ?", transactionID).Preload("TransactionDetails").First(&transaction).Error
	if err != nil {

		return transaction, err
	}

	return transaction, nil

}

func (r *repositoryTransactionImpl) FindAll(userID int) ([]Transaction, error) {

	var transactions []Transaction

	if userID != 0 {

		err := r.db.Model(&Transaction{}).Where("user_id = ?", userID).Preload("TransactionDetails").Find(&transactions).Error
		if err != nil {

			return transactions, err
		}

	} else {
		err := r.db.Model(&Transaction{}).Preload("TransactionDetails").Find(&transactions).Error
		if err != nil {

			return transactions, err
		}
	}

	return transactions, nil

}
