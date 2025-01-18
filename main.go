package main

import (
	"log"
	"project-akhir-awal/controller"
	"project-akhir-awal/middleware"
	"project-akhir-awal/repository"
	"project-akhir-awal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//!! Database
	//?? pindahin ke folder config nanti
	dsn := "root:@tcp(127.0.0.1:3306)/project_akhir_dibimbing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//!! Repository
	roleRepository := repository.NewRoleRepository(db)
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	//!! Service
	roleService := service.NewRoleService(roleRepository)
	userService := service.NewUserService(userRepository)
	productService := service.NewProductService(productRepository)
	transactionService := service.NewTransactionService(transactionRepository)
	authService := service.NewAuthService()

	//!! Middleware
	authMiddleware := middleware.AuthMiddleware(authService, userService)
	authMiddlewareAdmin := middleware.AuthMiddlewareAdmin(authService, userService)

	//!! Controller
	roleController := controller.NewRoleController(roleService)
	userController := controller.NewUserController(userService, authService)
	productController := controller.NewProductController(productService)
	transactionController := controller.NewTransactionController(transactionService)

	router := gin.Default()
	router.Static("/photo", "./public/user")
	router.Static("/image", "./public/product")
	api := router.Group("/api/v1")

	//!! Role Route
	api.POST("/role", roleController.CreateRole)
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? get role
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? update role
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? delete role

	//!! User Route
	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)
	api.POST("/email-checkers", userController.CheckEmailAvailability) //?? belum terealisasi ke register
	// api.POST("/upload_photo", authMiddleware(authService, userService), userController.UploadPhoto)
	api.POST("/upload-photo", authMiddleware, userController.UploadPhoto)
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? update user
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? delete user

	//!! Product Route
	api.GET("/products", productController.GetAllProducts)
	api.GET("/products/:id", productController.GetProduct)
	api.POST("/products", authMiddlewareAdmin, productController.CreateProduct)
	api.PUT("/products/:id", authMiddlewareAdmin, productController.UpdateProduct)
	api.POST("/product-image", authMiddlewareAdmin, productController.UploadImage)
	api.DELETE("/products/:id", authMiddlewareAdmin, productController.DeleteProduct)

	//!! Transaction Route
	api.GET("/products/:id/transactions", authMiddleware, transactionController.GetProductTransactions)
	api.GET("/transactions", authMiddleware, transactionController.GetUserTransactions)
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? create transaction
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? update transaction
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? pay transaction
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? delete transaction
	// api.POST("/upload-photo", authMiddleware, userController.UploadPhoto) //?? change status transaction

	router.Run()
}

// func authMiddleware(authService service.AuthService, userService service.UserService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenString)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)

// 		if !ok || !token.Valid {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		userID := int(claim["user_id"].(float64))

// 		user, err := userService.GetUserByID(userID)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		c.Set("currentUser", user)
// 	}
// }
