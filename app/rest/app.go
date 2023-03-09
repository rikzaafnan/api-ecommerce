package rest

import (
	"api-ecommerce/auth"
	common "api-ecommerce/common/uploader"
	"api-ecommerce/config"
	"api-ecommerce/database"
	"api-ecommerce/handler"
	"api-ecommerce/helper"
	"api-ecommerce/middleware"
	"api-ecommerce/payment"
	"api-ecommerce/product"
	"api-ecommerce/transaction"
	"api-ecommerce/user"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartApp() {

	db := database.InitializeDB()

	// route := gin.New()
	route := gin.Default()
	route.SetTrustedProxies(nil)
	route.Use(middleware.Logger())
	// route.Use(middleware.LoggingMiddlewareCustomTesting())

	route.Static("/images", "./upload-files/images")
	// route.Static("/uploads", "./upload-files")
	testRoute(route)
	userRoute(route, db)
	testRoutePakeJWT(route, db)
	productRoute(route, db)
	transactionRoute(route, db)
	paymentRoute(route, db)
	uploaderRoute(route, db)

	initScheduler()

	route.Run(config.SERVERPORT)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, "ping success")
	})

	route.MaxMultipartMemory = 8 << 20 // 8 MiB

	groupVersion := route.Group("/api/v1")
	groupVersion.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		module := c.PostForm("module")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {

			c.JSON(http.StatusBadRequest, err.Error())
			return

		}

		filename := filepath.Base(file.Filename)
		fileNameUPdate := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filename)
		pathDestination := "upload-files/images/" + fileNameUPdate

		if err := c.SaveUploadedFile(file, pathDestination); err != nil {

			c.JSON(http.StatusBadRequest, err.Error())
			return

		}

		// response := fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s and module=%s.", file.Filename, name, email, module)
		response := helper.APIResponse("success send picture", http.StatusOK, "success", map[string]string{
			"filename": file.Filename,
			"name":     name,
			"email":    email,
			"module":   module,
			"path":     pathDestination,
		})

		c.JSON(http.StatusOK, response)
	})
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
	routerGroupWithJWT.Use(middleware.UserLoginMiddleware())
	routerGroupWithJWT.GET("", transactionHandler.FindAll)
	routerGroupWithJWT.GET("/:transactionID", transactionHandler.FindByID)
	routerGroupWithJWT.POST("", transactionHandler.Create)

}

func paymentRoute(route *gin.Engine, db *gorm.DB) {

	authService := auth.NewService()
	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)

	transactionRepository := transaction.NewRepositoryTransaction(db)

	paymentRepository := payment.NewRepositoryPayment(db)
	paymentService := payment.NewServicePayment(paymentRepository, transactionRepository)
	paymentHandler := handler.Newpaymenthandler(paymentService)

	routerGroupWithJWT := route.Group("/users/:userID/payments")
	routerGroupWithJWT.Use(middleware.JWTMiddleware(authService, userService))
	routerGroupWithJWT.Use(middleware.UserLoginMiddleware())
	routerGroupWithJWT.GET("", paymentHandler.FindAll)
	routerGroupWithJWT.GET("/:paymentID", paymentHandler.FindByID)
	routerGroupWithJWT.POST("", paymentHandler.Create)

}

func uploaderRoute(route *gin.Engine, db *gorm.DB) {
	authService := auth.NewService()
	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)

	attachmentRepository := common.NewRepositoryAttachment(db)
	attachmentService := common.NewServiceAttachment(attachmentRepository)
	uploaderHandler := handler.NewUploaderhandler(attachmentService)

	routerGroupWithJWT := route.Group("/uploads")
	routerGroupWithJWT.Use(middleware.JWTMiddleware(authService, userService))
	routerGroupWithJWT.POST("", uploaderHandler.Save)
	routerGroupWithJWT.DELETE("/:uploaderID", uploaderHandler.Deleted)
}

// schedulerEvery5Minutes.StartAsync()
// schedulerEvery5Minutes := gocron
// _, err = schedulerEvery5Minutes.Every(5).Minute().StartAt(time.Now().UTC()).Do(func() {
// 	go kunyitTransactionServiceNoAggregate.TransactionStatusSyncToKINI(apps.RC)
// })
// if err != nil {
// 	log.Error(err)
// }
// schedulerEvery5Minutes

func initScheduler() {
	// testing cronjob
	log.Info("ini jam local : ", time.Now())
	schedulerEvery5Minutes := gocron.NewScheduler(time.Now().Location())
	_, err := schedulerEvery5Minutes.Every(5).Minute().StartAt(time.Now()).Do(func() {
		// go kunyitTransactionServiceNoAggregate.TransactionStatusSyncToKINI(apps.RC)
		log.Info("ini jalanin cronjob")
	})

	if err != nil {
		log.Error(err)
	}
	schedulerEvery5Minutes.StartAsync()
}
