package payment

import (
	"api-ecommerce/helper"
	"api-ecommerce/transaction"
	"api-ecommerce/user"
	"errors"
)

const (
	TransactionPaid = "paid"
)

type ServicePayment interface {
	Save(user user.User, input PaymentInput) (Payment, error)
	FindByID(PaymentID int) (Payment, error)
	FindAll(userID int) ([]Payment, error)
}

type servicePaymentImpl struct {
	repositoryPayment     RepositoryPayment
	repositoryTransaction transaction.RepositoryTransaction
}

func NewServicePayment(repositoryPayment RepositoryPayment, repositoryTransaction transaction.RepositoryTransaction) *servicePaymentImpl {
	return &servicePaymentImpl{repositoryPayment, repositoryTransaction}
}

func (s *servicePaymentImpl) Save(user user.User, input PaymentInput) (Payment, error) {
	var payment Payment

	transaction, err := s.repositoryTransaction.FindOneByUserIDAndCode(user.ID, input.TrxCode)
	if err != nil {
		return payment, errors.New("transaction not found")
	}

	if transaction.TransactionStatus == TransactionPaid {
		if err != nil {
			return payment, errors.New("transaction has been paid")
		}

	}

	payment.TransactionID = transaction.ID
	payment.InvoiceCode = helper.RandomString(10)
	payment.TotalAmount = transaction.TransactionTotal
	payment.Email = user.Email
	payment.Name = user.Name
	payment.UserID = user.ID

	payment, err = s.repositoryPayment.Save(payment)
	if err != nil {
		return payment, errors.New("payment error")
	}

	if payment.ID == 0 {
		return payment, errors.New("failed create payment")
	}

	payment, err = s.repositoryPayment.FindByID(payment.ID)
	if err != nil {
		return payment, err
	}

	return payment, nil

}

func (s *servicePaymentImpl) FindByID(PaymentID int) (Payment, error) {
	var payment Payment

	payment, err := s.repositoryPayment.FindByID(PaymentID)
	if err != nil {
		return payment, err
	}

	return payment, nil

}

func (s *servicePaymentImpl) FindAll(userID int) ([]Payment, error) {
	var payments []Payment

	payments, err := s.repositoryPayment.FindAll(userID)
	if err != nil {
		return payments, err
	}

	return payments, nil

}
