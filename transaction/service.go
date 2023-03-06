package transaction

import (
	"api-ecommerce/helper"
	"api-ecommerce/product"
	"errors"
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

	var totalTransaction = 0
	var transactionDetails []TransactionDetail

	for index, productDetail := range input.ProductDetails {
		var totalAmount = 0

		var transactionDetail TransactionDetail

		var product product.Product

		product, err := s.repositoryProduct.FindByID(input.ProductDetails[index].ProductID)
		if err != nil {
			return transaction, errors.New("product not found")
		}

		totalAmount = productDetail.Quantity * product.Amount
		transactionDetail.ProductID = product.ID
		transactionDetail.Quantity = productDetail.Quantity
		transactionDetail.ProductName = product.Name
		transactionDetail.ProductPrice = product.Amount
		transactionDetail.TotalAmount = totalAmount

		totalTransaction = totalTransaction + totalAmount
		transactionDetails = append(transactionDetails, transactionDetail)

	}

	transaction.UserID = userID
	transaction.Name = input.Name
	transaction.Email = input.Email
	transaction.Code = helper.RandomString(10)
	transaction.Address = input.Address
	transaction.TransactionStatus = "pending"
	transaction.TransactionTotal = totalTransaction
	transaction.TransactionDetails = transactionDetails

	transaction, err := s.repositoryTransaction.Save(transaction)
	if err != nil {
		return transaction, err
	}

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
