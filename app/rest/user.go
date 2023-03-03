package rest

import (
	"api-ecommerce/helper"
	"api-ecommerce/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.ServiceUser
}

func NewUserhandler(userService user.ServiceUser) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterInput
	// struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		// errors := helper.FormatValidationError(err)

		// errorMessage := gin.H{"errors": errors}

		// response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", "error message")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(newUser, "isinya token")

	response := helper.APIResponse("account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.Logininput

	// menagmbil data inputan dari user lalu bind ke struct input login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// errors := helper.FormatValidationError(err)

		// errorMessage := gin.H{"errors": errors}

		// response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", "errorMessage")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "token")
	response := helper.APIResponse("Succesfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
