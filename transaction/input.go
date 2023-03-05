package transaction

type TransactionInput struct {
	Name            string          `json:"name" binding:"required"`
	Address         string          `json:"address" binding:"required"`
	TrasactionTotal int             `json:"transactionTotal" binding:"required"`
	ProductDetails  []ProductDetail `json:"productDetails"`
}

type ProductDetail struct {
	ProductID int `json:"productID" `
	Quantity  int `json:"quantity" `
}
