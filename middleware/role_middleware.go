package middleware

import (
	"api-ecommerce/helper"
	"api-ecommerce/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {

		currentUser := c.MustGet("currentUser").(user.User)

		next := false

		for _, role := range roles {

			if currentUser.Role == role {
				next = true
			}

		}

		if !next {
			response := helper.APIResponse("user not permission", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}

}

func UserLoginMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		currentUser := c.MustGet("currentUser").(user.User)

		paramUserID := c.Param("userID")
		userID, _ := strconv.Atoi(paramUserID)

		if currentUser.ID != userID {
			response := helper.APIResponse("user not permission", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}

}
