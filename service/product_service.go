package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"
)

type ProductService interface {
	FindAllProducts() ([]entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) FindAllProducts() ([]entity.Product, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}
