package controller

import (
	"net/http"
	"project-akhir-awal/helper"
	"project-akhir-awal/service"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	roleService service.RoleService
}

func NewRoleController(roleService service.RoleService) *roleController {
	return &roleController{roleService}
}

func (h *roleController) CreateRole(c *gin.Context) {
	var input service.CreateRoleInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create Role Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRole, err := h.roleService.CreateRole(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create Role Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Role has been registered", http.StatusOK, "success", newRole)

	c.JSON(http.StatusOK, response)
}
