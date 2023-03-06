package transaction

type TransactionFormatter struct {
	ID                 int                          `json:"id"`
	UserID             int                          `json:"useriID"`
	Name               string                       `json:"name"`
	Email              string                       `json:"email"`
	Code               string                       `json:"code"`
	Address            string                       `json:"address"`
	TransactionStatus  string                       `json:"transactionStatus"`
	TransactionTotal   int                          `json:"transactionTotal"`
	Description        string                       `json:"description"`
	TransactionDetails []TransactionDetailFormatter `json:"transactionDetails"`
}

func FormatTransaction(transaction Transaction, transactionDetail []TransactionDetail) TransactionFormatter {

	var TransactionDetailFormatters []TransactionDetailFormatter

	for _, transactionDetailValue := range transactionDetail {

		transactionDetailFormatter := FormatTransactionDetail(transactionDetailValue)
		TransactionDetailFormatters = append(TransactionDetailFormatters, transactionDetailFormatter)

	}

	formatter := TransactionFormatter{
		ID:                 transaction.ID,
		UserID:             transaction.UserID,
		Name:               transaction.Name,
		Email:              transaction.Email,
		Code:               transaction.Code,
		Address:            transaction.Address,
		TransactionStatus:  transaction.TransactionStatus,
		TransactionTotal:   transaction.TransactionTotal,
		Description:        transaction.Description,
		TransactionDetails: TransactionDetailFormatters,
	}

	return formatter
}

type TransactionDetailFormatter struct {
	ID           int    `json:"id"`
	ProductID    int    `json:"productId"`
	ProductName  string `json:"productName"`
	Quantity     int    `json:"quantity"`
	ProductPrice int    `json:"productPrice"`
	TotalAmount  int    `json:"totalAmount"`
}

func FormatTransactionDetail(transactionDetail TransactionDetail) TransactionDetailFormatter {
	formatter := TransactionDetailFormatter{
		ID:           transactionDetail.ID,
		ProductID:    transactionDetail.ProductID,
		ProductName:  transactionDetail.ProductName,
		Quantity:     transactionDetail.Quantity,
		ProductPrice: transactionDetail.ProductPrice,
		TotalAmount:  transactionDetail.TotalAmount,
	}

	return formatter
}
