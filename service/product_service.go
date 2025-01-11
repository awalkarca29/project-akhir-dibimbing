package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
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
