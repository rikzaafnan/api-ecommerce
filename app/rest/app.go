package rest

import (
	"api-ecommerce/config"
	"api-ecommerce/database"
	"api-ecommerce/user"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func StartApp() {

	database.InitializeDB()

	db := database.GetDB()

	log.Info(db)

	route := gin.Default()

	testRoute(route)
	userRoute(route, db)

	route.Run(config.SERVERPORT)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}

func userRoute(route *gin.Engine, db *gorm.DB) {

	userRepository := user.NewRepositoryUser(db)
	userService := user.NewServiceUser(userRepository)

	userHandler := NewUserhandler(userService)

	// no jwt
	routeGroup := route.Group("/users")
	routeGroup.POST("/register", userHandler.RegisterUser)
	routeGroup.POST("/login", userHandler.Login)

	// routerGroupWithJWT := route.Group("/users")
	// routerGroupWithJWT.Use(authService.Authentication())
	// routerGroupWithJWT.PUT("/:userID", userHandler.Update)
	// routerGroupWithJWT.DELETE("/:userID", userHandler.Delete)
	// routerGroupWithJWT.GET("/me", userHandler.Me)
	// routerGroupWithJWT.PUT("/verification", userHandler.Me)

}
