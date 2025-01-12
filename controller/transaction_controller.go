package controller

import (
	"net/http"
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

func (h *transactionController) GetCampaignTransaction(c *gin.Context) {
	var input service.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get product's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetTransactionByProductID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get product's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product's transaction detail", http.StatusOK, "success", helper.FormatProductTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
