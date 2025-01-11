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

type ProductDetailFormatter struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Slug        string                  `json:"slug"`
	ImageURL    string                  `json:"image_url"`
	Price       int                     `json:"price"`
	Stock       int                     `json:"int"`
	Images      []ProductImageFormatter `json:"images"`
}

type ProductImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatProductDetail(product entity.Product) ProductDetailFormatter {
	productDetailFormatter := ProductDetailFormatter{}
	productDetailFormatter.ID = product.ID
	productDetailFormatter.Name = product.Name
	productDetailFormatter.Description = product.Description
	productDetailFormatter.Slug = product.Slug
	productDetailFormatter.Price = product.Price
	productDetailFormatter.Stock = product.Stock
	productDetailFormatter.ImageURL = ""

	if len(product.ProductImages) > 0 {
		productDetailFormatter.ImageURL = product.ProductImages[0].FileName
	}

	images := []ProductImageFormatter{}

	for _, image := range product.ProductImages {
		productImageFormatter := ProductImageFormatter{}
		productImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		productImageFormatter.IsPrimary = isPrimary

		images = append(images, productImageFormatter)
	}

	productDetailFormatter.Images = images

	return productDetailFormatter
}
