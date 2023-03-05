package rest

import (
	"api-ecommerce/auth"
	"api-ecommerce/config"
	"api-ecommerce/database"
	"api-ecommerce/handler"
	"api-ecommerce/middleware"
	"api-ecommerce/product"
	"api-ecommerce/transaction"
	"api-ecommerce/user"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func StartApp() {

	db := database.InitializeDB()

	log.Info(db)

	route := gin.Default()

	testRoute(route)
	userRoute(route, db)
	testRoutePakeJWT(route, db)
	productRoute(route, db)
	transactionRoute(route, db)

	route.Run(config.SERVERPORT)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}

func testRoutePakeJWT(route *gin.Engine, db *gorm.DB) {

	authService := auth.NewService()

	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)
	userHandler := handler.NewUserhandler(userService, authService)

	routerGroupWithJWT := route.Group("/test")
	routerGroupWithJWT.Use(middleware.JWTMiddleware(authService, userService))
	routerGroupWithJWT.GET("/with-jwt", userHandler.TestJWT)
}

func userRoute(route *gin.Engine, db *gorm.DB) {

	authService := auth.NewService()
	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)

	userHandler := handler.NewUserhandler(userService, authService)

	// no jwt
	routeGroup := route.Group("/users")
	// bypass for refgister admin
	routeGroup.POST("/register-by-pass", userHandler.RegisterByPassSuperUser)

	routeGroup.POST("/register", userHandler.RegisterUser)
	routeGroup.POST("/login", userHandler.Login)
	routeGroup.GET("/verification", userHandler.PatchVerification)

	routerGroupWithJWT := route.Group("/users")
	routerGroupWithJWT.Use(middleware.JWTMiddleware(authService, userService))
	// routerGroupWithJWT.PUT("/:userID", userHandler.Update)
	// routerGroupWithJWT.DELETE("/:userID", userHandler.Delete)
	routerGroupWithJWT.GET("/me", userHandler.Me)

}

func productRoute(route *gin.Engine, db *gorm.DB) {

	authService := auth.NewService()
	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)

	productRepository := product.NewRepositoryProduct(db)
	productService := product.NewServiceProduct(productRepository)
	productHandler := handler.NewProducthandler(productService)

	routerGroupWithJWT := route.Group("/products")
	routerGroupWithJWT.Use(middleware.JWTMiddleware(authService, userService))
	routerGroupWithJWT.GET("", productHandler.FindAll)
	routerGroupWithJWT.GET("/:productID", productHandler.FindByID)
	routerGroupWithJWT.Use(middleware.RoleMiddleware("admin"))
	routerGroupWithJWT.POST("", productHandler.Create)
	routerGroupWithJWT.PUT("/:productID", productHandler.Update)
	routerGroupWithJWT.DELETE("/:productID", productHandler.Delete)

}

func transactionRoute(route *gin.Engine, db *gorm.DB) {

	authService := auth.NewService()
	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)

	productRepository := product.NewRepositoryProduct(db)

	transactionRepository := transaction.NewRepositoryTransaction(db)
	transactiontService := transaction.NewServiceTransaction(transactionRepository, productRepository)
	transactionHandler := handler.NewTransactionhandler(transactiontService)

	routerGroupWithJWT := route.Group("/users/:userID/transactions")
	routerGroupWithJWT.Use(middleware.JWTMiddleware(authService, userService))
	routerGroupWithJWT.GET("", transactionHandler.FindAll)
	routerGroupWithJWT.GET("/:transactionID", transactionHandler.FindByID)
	routerGroupWithJWT.POST("", transactionHandler.Create)

}
