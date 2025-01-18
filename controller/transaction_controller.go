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

func (h *transactionController) GetTransaction(c *gin.Context) {
	var input service.GetTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionDetail, err := h.transactionService.GetTransactionByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Transaction detail", http.StatusOK, "success", helper.FormatCreateTransaction(transactionDetail))
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
		response := helper.APIResponse("Failed to pay transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	markPaid, err := h.transactionService.MarkPaid(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to pay transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to pay transaction", http.StatusOK, "success", helper.FormatCreateTransaction(markPaid))
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) MarkCancel(c *gin.Context) {
	var inputID service.GetTransactionInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to cancel transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	markCancel, err := h.transactionService.MarkCancel(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to cancel transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to cancel transaction", http.StatusOK, "success", helper.FormatCreateTransaction(markCancel))
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) MarkStatus(c *gin.Context) {
	var inputID service.GetTransactionInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to change status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData service.GetTransactionStatusInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to change status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	markStatus, err := h.transactionService.MarkStatus(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to change status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to change status", http.StatusOK, "success", helper.FormatCreateTransaction(markStatus))
	c.JSON(http.StatusOK, response)
}
