package middleware

import (
	"api-ecommerce/helper"
	"api-ecommerce/user"
	"fmt"
	"net/http"

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

		// authHeader := c.GetHeader("Authorization")

		// if !strings.Contains(authHeader, "Bearer") {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// tokenString := ""
		// // untuk memisahkan antara bearer dan token
		// arraytoken := strings.Split(authHeader, " ")
		// if len(arraytoken) == 2 {
		// 	tokenString = arraytoken[1]
		// }

		// token, err := authService.ValidateToken(tokenString)
		// if err != nil {
		// 	response := helper.APIResponse(err.Error(), http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// claim, ok := token.Claims.(jwt.MapClaims)
		// if !ok || !token.Valid {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// userID := int(claim["user_id"].(float64))

		// user, err := userService.GetUserByID(userID)
		// if err != nil {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// // menyimpan data di set untuk bisa dipakai di file manapun
		// c.Set("currentUser", user)

		fmt.Println("melewati role middleware")

	}

}
