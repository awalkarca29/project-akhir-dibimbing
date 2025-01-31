package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"project-akhir-awal/helper"
	"project-akhir-awal/service"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *productController {
	return &productController{productService}
}

func (h *productController) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", helper.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productController) GetProduct(c *gin.Context) {
	var input service.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.productService.GetProductByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product detail", http.StatusOK, "success", helper.FormatProductDetail(productDetail))
	c.JSON(http.StatusOK, response)
}

func (h *productController) CreateProduct(c *gin.Context) {
	var input service.CreateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create product", http.StatusOK, "success", helper.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productController) UpdateProduct(c *gin.Context) {
	var inputID service.GetProductDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData service.CreateProductInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedProduct, err := h.productService.UpdateProduct(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update product", http.StatusOK, "success", helper.FormatProduct(updatedProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productController) UploadImage(c *gin.Context) {
	var input service.UploadImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// currentProduct := c.MustGet("currentUser").(entity.User)
	// userID := currentUser.ID

	path := fmt.Sprintf("public/product/%d-%s", rand.Int(), file.Filename)
	// path := fmt.Sprintf("public/product/%d-%s", input.ProductID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.UploadImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Image successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *productController) DeleteProduct(c *gin.Context) {
	var input service.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedProduct, err := h.productService.DeleteProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete product", http.StatusOK, "success", helper.FormatProductDetail(deletedProduct))
	c.JSON(http.StatusOK, response)
}
