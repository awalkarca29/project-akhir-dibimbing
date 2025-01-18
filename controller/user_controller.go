package controller

import (
	"fmt"
	"net/http"
	"project-akhir-awal/entity"
	"project-akhir-awal/helper"
	"project-akhir-awal/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserController(userService service.UserService, authService service.AuthService) *userController {
	return &userController{userService, authService}
}

func (h *userController) Register(c *gin.Context) {
	var input service.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID, newUser.RoleID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := helper.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var input service.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID, loginUser.RoleID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := helper.FormatUser(loginUser, token)

	response := helper.APIResponse("Successfully Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) CheckEmailAvailability(c *gin.Context) {
	var input service.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	path := fmt.Sprintf("public/user/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UploadPhoto(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Photo successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
	var inputData service.UserInput

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	updatedUser, err := h.userService.UpdateUser(userID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update user", http.StatusOK, "success", helper.FormatUserUpdate(updatedUser))
	c.JSON(http.StatusOK, response)
}

func (h *userController) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	deletedUser, err := h.userService.DeleteUser(userID)
	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete user", http.StatusOK, "success", helper.FormatUserUpdate(deletedUser))
	c.JSON(http.StatusOK, response)
}

func (h *userController) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		response := helper.APIResponse("Error to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of users", http.StatusOK, "success", helper.FormatUsers(users))
	c.JSON(http.StatusOK, response)
}
