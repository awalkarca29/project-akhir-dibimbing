package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"

	"github.com/gosimple/slug"
)

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(input GetProductDetailInput) (entity.Product, error)
	CreateProduct(input CreateProductInput) (entity.Product, error)
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

func (s *productService) CreateProduct(input CreateProductInput) (entity.Product, error) {
	product := entity.Product{}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.Slug = slug.Make(input.Name)

	newProduct, err := s.productRepository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}
