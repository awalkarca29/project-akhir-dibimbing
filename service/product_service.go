package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"
)

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(input GetProductDetailInput) (entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *productService) GetProductByID(input GetProductDetailInput) (entity.Product, error) {
	product, err := s.productRepository.FindByID(input.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}
