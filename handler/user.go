package handler

import (
	"api-ecommerce/auth"
	"api-ecommerce/helper"
	"api-ecommerce/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.ServiceUser
	authService auth.ServiceAuth
}

func NewUserhandler(userService user.ServiceUser, authService auth.ServiceAuth) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	token, err := h.authService.GenerateToken(newUser.ID, newUser.Role)
	if err != nil {

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
func (h *userHandler) RegisterByPassSuperUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUserByPassSuperAdmin(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	token, err := h.authService.GenerateToken(newUser.ID, newUser.Role)
	if err != nil {

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.Logininput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		// response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", "errorMessage")
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

	token, err := h.authService.GenerateToken(loggedinUser.ID, loggedinUser.Role)
	if err != nil {

		response := helper.APIResponse("generate account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Succesfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) PatchVerification(c *gin.Context) {

	email := c.Query("email")

	loggedinUser, err := h.userService.Verification(email)
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

func (h *userHandler) Me(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.APIResponse("Succesfully get Me", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) TestJWT(c *gin.Context) {

	email := c.Query("email")
	fmt.Println(email)

	// loggedinUser, err := h.userService.Verification(email)
	// if err != nil {

	// 	errorMessage := gin.H{"errors": err.Error()}
	// 	response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// formatter := user.FormatUser(loggedinUser, "token")
	// response := helper.APIResponse("Succesfully loggedin", http.StatusOK, "success", formatter)

	user := c.MustGet("currentUser").(user.User)

	c.JSON(http.StatusOK, user)

}
