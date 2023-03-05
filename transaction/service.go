package transaction

import (
	"api-ecommerce/helper"
	"api-ecommerce/product"
	"errors"
	"fmt"
)

type ServiceTransaction interface {
	Save(userID int, input TransactionInput) (Transaction, error)
	FindByID(transactionID int) (Transaction, error)
	FindAll(userID int) ([]Transaction, error)
}

type serviceTransactionImpl struct {
	repositoryTransaction RepositoryTransaction
	repositoryProduct     product.RepositoryProduct
}

func NewServiceTransaction(repositoryTransaction RepositoryTransaction, repositoryProduct product.RepositoryProduct) *serviceTransactionImpl {
	return &serviceTransactionImpl{repositoryTransaction, repositoryProduct}
}

func (s *serviceTransactionImpl) Save(userID int, input TransactionInput) (Transaction, error) {
	var transaction Transaction

	fmt.Println("ini service create transaction")

	var totalAmount = 0
	var transactionDetails []TransactionDetail

	for _, productDetail := range input.ProductDetails {

		var transactionDetail TransactionDetail

		product, err := s.repositoryProduct.FindByID(productDetail.ProductID)
		if err != nil {
			return transaction, errors.New("product not found")
		}

		totalAmount = productDetail.Quantity * product.Amount
		transactionDetail.ProductID = product.ID

		transactionDetails = append(transactionDetails, transactionDetail)

	}

	transaction.Name = input.Name
	transaction.Code = helper.RandomString(10)
	transaction.Address = input.Address
	transaction.TransactionStatus = "pending"
	transaction.TransactionTotal = totalAmount
	transaction.TransactionDetails = transactionDetails

	fmt.Println(transactionDetails)

	// transaction, err := s.repositoryTransaction.Save(transaction)
	// if err != nil {
	// 	return transaction, err
	// }

	return transaction, nil

}

func (s *serviceTransactionImpl) FindByID(transactionID int) (Transaction, error) {
	var transaction Transaction

	transaction, err := s.repositoryTransaction.FindByID(transactionID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (s *serviceTransactionImpl) FindAll(userID int) ([]Transaction, error) {
	var transactions []Transaction

	transactions, err := s.repositoryTransaction.FindAll(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil

}
