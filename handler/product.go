package handler

import (
	"api-ecommerce/helper"
	"api-ecommerce/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.ServiceProduct
}

func NewProducthandler(productService product.ServiceProduct) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) Create(c *gin.Context) {

	var input product.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("create product failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := h.productService.Create(input)
	if err != nil {

		response := helper.APIResponse("create product failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := product.FormatProduct(newProduct)

	response := helper.APIResponse("product has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *productHandler) Update(c *gin.Context) {

	var input product.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("update product failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	paramProductID := c.Param("productID")
	productID, _ := strconv.Atoi(paramProductID)

	updateProduct, err := h.productService.Update(productID, input)
	if err != nil {

		response := helper.APIResponse("update product failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := product.FormatProduct(updateProduct)

	response := helper.APIResponse("product has been updatede", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *productHandler) Delete(c *gin.Context) {

	paramProductID := c.Param("productID")
	productID, _ := strconv.Atoi(paramProductID)

	err := h.productService.Delete(productID)
	if err != nil {

		response := helper.APIResponse("create product failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.APIResponse("product has been deleted", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)

}

func (h *productHandler) FindByID(c *gin.Context) {

	paramProductID := c.Param("productID")
	productID, _ := strconv.Atoi(paramProductID)

	productData, err := h.productService.FindByID(productID)
	if err != nil {

		response := helper.APIResponse("find product failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := product.FormatProduct(productData)

	response := helper.APIResponse("product has been updatede", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *productHandler) FindAll(c *gin.Context) {

	var productFormatters []product.ProductFormatter

	productDatas, err := h.productService.FindAll()
	if err != nil {

		response := helper.APIResponse("get product failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	if len(productDatas) > 0 {
		for _, productValue := range productDatas {
			formatter := product.FormatProduct(productValue)
			productFormatters = append(productFormatters, formatter)

		}
	}

	response := helper.APIResponse("success get product", http.StatusOK, "success", productFormatters)

	c.JSON(http.StatusOK, response)

}
