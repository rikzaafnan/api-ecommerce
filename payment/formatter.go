package payment

import "api-ecommerce/transaction"

type PaymentFormatter struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transactionID"`
	InvoiceCode   string `json:"invoiceCode"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	TotalAmount   int    `json:"totalAmount"`
	Transaction   transaction.TransactionFormatter
}

func FormatterPayment(payment Payment) PaymentFormatter {

	transactionformatter := transaction.FormatTransaction(payment.Transaction, payment.Transaction.TransactionDetails)

	formatter := PaymentFormatter{
		ID:            payment.ID,
		TransactionID: payment.TransactionID,
		InvoiceCode:   payment.InvoiceCode,
		Email:         payment.Email,
		Name:          payment.Name,
		TotalAmount:   payment.TotalAmount,
		Transaction:   transactionformatter,
	}

	return formatter
}
