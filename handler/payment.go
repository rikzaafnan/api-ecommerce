package handler

import (
	"api-ecommerce/helper"
	"api-ecommerce/payment"
	"api-ecommerce/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.ServicePayment
}

func Newpaymenthandler(paymentService payment.ServicePayment) *paymentHandler {
	return &paymentHandler{paymentService}
}

func (h *paymentHandler) Create(c *gin.Context) {

	var input payment.PaymentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("create transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	newPayment, err := h.paymentService.Save(currentUser, input)
	if err != nil {

		response := helper.APIResponse("create payment failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := payment.FormatterPayment(newPayment)

	response := helper.APIResponse("payment has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *paymentHandler) FindByID(c *gin.Context) {

	paramPaymentID := c.Param("paymentID")
	paymentID, _ := strconv.Atoi(paramPaymentID)

	paymentData, err := h.paymentService.FindByID(paymentID)
	if err != nil {

		response := helper.APIResponse("find transaction failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := payment.FormatterPayment(paymentData)

	response := helper.APIResponse("success find payment", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *paymentHandler) FindAll(c *gin.Context) {

	var paymentFormatters []payment.PaymentFormatter

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

	payments, err := h.paymentService.FindAll(userIDQueryString)
	if err != nil {

		response := helper.APIResponse("get payment failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	if len(payments) > 0 {
		for _, paymentData := range payments {
			formatter := payment.FormatterPayment(paymentData)
			paymentFormatters = append(paymentFormatters, formatter)

		}
	}

	response := helper.APIResponse("success get payment", http.StatusOK, "success", paymentFormatters)

	c.JSON(http.StatusOK, response)

}
