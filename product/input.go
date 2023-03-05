package product

type ProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Amount      int    `json:"amount" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
