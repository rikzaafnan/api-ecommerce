package product

type ProductFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Quantity    int    `json:"quantity"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		Amount:      product.Amount,
		Quantity:    product.Quantity,
	}

	return formatter
}
