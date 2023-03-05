package product

import "gorm.io/gorm"

type RepositoryProduct interface {
	Save(product Product) (Product, error)
	FindByID(ID int) (Product, error)
	DeleteByID(ID int) error
	Update(productID int, product Product) (Product, error)
	FindAllProduct() ([]Product, error)
}

type repositoryProductImpl struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProductImpl {
	return &repositoryProductImpl{db}
}

func (r *repositoryProductImpl) Save(product Product) (Product, error) {

	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryProductImpl) FindByID(ID int) (Product, error) {
	var product Product

	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {

		return product, err
	}

	return product, nil
}

func (r *repositoryProductImpl) Update(productID int, product Product) (Product, error) {

	var productModel Product

	err := r.db.Where("id = ?", productID).First(&productModel).Error
	if err != nil {

		return product, err
	}

	err = r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryProductImpl) DeleteByID(ID int) error {
	var product Product

	err := r.db.Where("id = ?", ID).Delete(&product).Error
	if err != nil {

		return err
	}

	return nil
}

func (r *repositoryProductImpl) FindAllProduct() ([]Product, error) {
	var products []Product

	err := r.db.Find(&products).Error
	if err != nil {

		return products, err
	}

	return products, nil
}
