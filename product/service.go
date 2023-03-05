package product

import (
	"errors"
	"fmt"
)

type ServiceProduct interface {
	Create(input ProductInput) (Product, error)
	FindByID(productID int) (Product, error)
	Update(productID int, input ProductInput) (Product, error)
	Delete(productID int) error
	FindAll() ([]Product, error)
}

type serviceProductImpl struct {
	repositoryProduct RepositoryProduct
}

func NewServiceProduct(repositoryProduct RepositoryProduct) *serviceProductImpl {
	return &serviceProductImpl{repositoryProduct}
}

func (s *serviceProductImpl) Create(input ProductInput) (Product, error) {
	var product Product

	product.Name = input.Name
	product.Description = input.Description
	product.Amount = input.Amount
	product.Quantity = input.Quantity

	product, err := s.repositoryProduct.Save(product)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("failed create product")
	}

	return product, nil

}

func (s *serviceProductImpl) FindByID(productID int) (Product, error) {
	var product Product

	product, err := s.repositoryProduct.FindByID(productID)
	if err != nil {
		fmt.Println(err)
		err = errors.New("product not found")
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("failed find product")
	}

	return product, nil

}

func (s *serviceProductImpl) Update(productID int, input ProductInput) (Product, error) {
	var product Product

	product, err := s.repositoryProduct.FindByID(productID)
	if err != nil {
		return product, err
	}

	fmt.Println("kesini kah?")

	product.Name = input.Name
	product.Description = input.Description
	product.Amount = input.Amount
	product.Quantity = input.Quantity

	product, err = s.repositoryProduct.Update(productID, product)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("failed update product")
	}

	product, err = s.repositoryProduct.FindByID(productID)
	if err != nil {
		return product, err
	}

	return product, nil

}

func (s *serviceProductImpl) Delete(productID int) error {

	var product Product

	product, err := s.repositoryProduct.FindByID(productID)
	if err != nil {
		fmt.Println(err)
		err = errors.New("product not found")
		return err
	}

	if product.ID == 0 {
		return errors.New("failed find product")
	}

	err = s.repositoryProduct.DeleteByID(productID)
	if err != nil {
		return err
	}

	return nil

}

func (s *serviceProductImpl) FindAll() ([]Product, error) {

	var products []Product

	products, err := s.repositoryProduct.FindAllProduct()
	if err != nil {
		return products, err
	}

	return products, nil

}
