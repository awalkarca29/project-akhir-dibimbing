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
	dsn := "root:@tcp(127.0.0.1:3306)/project_akhir_dibimbing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	roleRepository := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepository)
	roleController := controller.NewRoleController(roleService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	authMiddleware := middleware.AuthMiddleware(authService, userService)
	userController := controller.NewUserController(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/role", roleController.CreateRole)

	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	// api.POST("/upload_photo", authMiddleware(authService, userService), userController.UploadPhoto)
	api.POST("/upload_photo", authMiddleware, userController.UploadPhoto)

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
