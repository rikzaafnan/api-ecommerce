package handler

import (
	"api-ecommerce/helper"
	"api-ecommerce/transaction"
	"api-ecommerce/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactiontHandler struct {
	transactionService transaction.ServiceTransaction
}

func NewTransactionhandler(transactionService transaction.ServiceTransaction) *transactiontHandler {
	return &transactiontHandler{transactionService}
}

func (h *transactiontHandler) Create(c *gin.Context) {

	var input transaction.TransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("create transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	newTransaction, err := h.transactionService.Save(currentUser.ID, input)
	if err != nil {

		response := helper.APIResponse("create transaction failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := transaction.FormatTransaction(newTransaction, newTransaction.TransactionDetails)

	response := helper.APIResponse("transaction has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *transactiontHandler) FindByID(c *gin.Context) {

	paramTransactionID := c.Param("transactionID")
	transactionID, _ := strconv.Atoi(paramTransactionID)

	transactionData, err := h.transactionService.FindByID(transactionID)
	if err != nil {

		response := helper.APIResponse("find transaction failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}
	formatter := transaction.FormatTransaction(transactionData, transactionData.TransactionDetails)

	response := helper.APIResponse("success find transaction", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *transactiontHandler) FindAll(c *gin.Context) {

	var transactionFormatters []transaction.TransactionFormatter

	var userIDQueryString int

	queryParamUserID := c.Query("user_id")

	if queryParamUserID == "" {

		userIDQueryString = 0

	} else {
		userID, err := strconv.Atoi(queryParamUserID)
		if err != nil {
			response := helper.APIResponse("error parsing user_id", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userIDQueryString = userID
	}

	transactions, err := h.transactionService.FindAll(userIDQueryString)
	if err != nil {

		response := helper.APIResponse("get product failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	if len(transactions) > 0 {
		for _, transactionData := range transactions {
			formatter := transaction.FormatTransaction(transactionData, transactionData.TransactionDetails)
			transactionFormatters = append(transactionFormatters, formatter)

		}
	}

	response := helper.APIResponse("success get product", http.StatusOK, "success", transactionFormatters)

	c.JSON(http.StatusOK, response)

}
