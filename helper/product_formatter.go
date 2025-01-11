package helper

import "project-akhir-awal/entity"

type ProductFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

func FormatProduct(product entity.Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.Name = product.Name
	productFormatter.Description = product.Description
	productFormatter.Slug = product.Slug
	productFormatter.Price = product.Price
	productFormatter.Stock = product.Stock
	productFormatter.ImageURL = ""

	if len(product.ProductImages) > 0 {
		productFormatter.ImageURL = product.ProductImages[0].FileName
	}

	return productFormatter
}

func FormatProducts(products []entity.Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
