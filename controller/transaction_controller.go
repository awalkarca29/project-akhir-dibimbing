package controller

import (
	"net/http"
	"project-akhir-awal/entity"
	"project-akhir-awal/helper"
	"project-akhir-awal/service"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *transactionController {
	return &transactionController{transactionService}
}

func (h *transactionController) GetProductTransactions(c *gin.Context) {
	var input service.GetProductTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get product's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetTransactionByProductID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get product's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product's transactions", http.StatusOK, "success", helper.FormatProductTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	transactions, err := h.transactionService.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", helper.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) CreateTransaction(c *gin.Context) {
	var input service.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser
	// input.Product = inputID

	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create transaction", http.StatusOK, "success", helper.FormatCreateTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) MarkPaid(c *gin.Context) {
	var inputID service.GetTransactionInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	markPaid, err := h.transactionService.MarkPaid(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update status", http.StatusOK, "success", helper.FormatCreateTransaction(markPaid))
	c.JSON(http.StatusOK, response)
}
