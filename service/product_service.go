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

type UploadImageInput struct {
	ProductID int  `form:"product_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
}

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(input GetProductDetailInput) (entity.Product, error)
	CreateProduct(input CreateProductInput) (entity.Product, error)
	UpdateProduct(inputID GetProductDetailInput, inputData CreateProductInput) (entity.Product, error)
	DeleteProduct(inputID GetProductDetailInput) (entity.Product, error)
	UploadImage(input UploadImageInput, fileLocation string) (entity.ProductImage, error)
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

func (s *productService) UpdateProduct(inputID GetProductDetailInput, inputData CreateProductInput) (entity.Product, error) {
	product, err := s.productRepository.FindByID(inputID.ID)
	if err != nil {
		return product, err
	}

	product.Name = inputData.Name
	product.Description = inputData.Description
	product.Price = inputData.Price
	product.Stock = inputData.Stock
	// product.Slug = slug.Make(inputData.Name)

	updatedProduct, err := s.productRepository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *productService) UploadImage(input UploadImageInput, fileLocation string) (entity.ProductImage, error) {
	isPrimary := 0

	if input.IsPrimary {
		isPrimary = 1

		_, err := s.productRepository.MarkAllImagesAsNonPrimary(input.ProductID)
		if err != nil {
			return entity.ProductImage{}, err
		}
	}

	productImage := entity.ProductImage{}
	productImage.ProductID = input.ProductID
	productImage.IsPrimary = isPrimary
	productImage.FileName = fileLocation

	newUploadImage, err := s.productRepository.UploadImage(productImage)
	if err != nil {
		return newUploadImage, err
	}

	return newUploadImage, nil
}

func (s *productService) DeleteProduct(inputID GetProductDetailInput) (entity.Product, error) {
	product, err := s.productRepository.FindByID(inputID.ID)
	if err != nil {
		return product, err
	}

	deleteProduct, err := s.productRepository.Delete(product)
	if err != nil {
		return deleteProduct, err
	}

	return deleteProduct, nil
}
