package rest

import (
	"api-ecommerce/config"
	"api-ecommerce/database"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func StartApp() {

	database.InitializeDB()

	db := database.GetDB()

	log.Info(db)

	route := gin.Default()

	testRoute(route)
	// userRoute(route, db)

	route.Run(config.SERVERPORT)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}

func userRoute(route *gin.Engine) {

	// userRepository := userpg.NewUserPG(db)
	// userService := service.NewUserService(userRepository)
	// userHandler := NewUserhandler(userService)
	// authService := service.NewAuthService(userRepository)

	// // no jwt
	// routeGroup := route.Group("/users")

	// routeGroup.POST("/register", userHandler.Register)
	// routeGroup.POST("/login", userHandler.Login)

	// routerGroupWithJWT := route.Group("/users")
	// routerGroupWithJWT.Use(authService.Authentication())
	// routerGroupWithJWT.PUT("/:userID", userHandler.Update)
	// routerGroupWithJWT.DELETE("/:userID", userHandler.Delete)
	// routerGroupWithJWT.GET("/me", userHandler.Me)
}
