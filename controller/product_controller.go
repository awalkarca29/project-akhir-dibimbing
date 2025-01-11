package controller

import (
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
