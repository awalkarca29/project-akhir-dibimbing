package main

import (
	"log"
	"project-akhir-awal/controller"
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
	userController := controller.NewUserController(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/role", roleController.CreateRole)

	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/upload_photo", userController.UploadPhoto)

	router.Run()
}
