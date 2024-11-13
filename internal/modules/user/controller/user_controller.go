package controller

import (
	"github.com/gin-gonic/gin"
	"goproject/internal/modules/user/dto"
	"goproject/internal/modules/user/service"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(sm *service.UserService) *UserController {
	return &UserController{userService: sm}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userService.CreateUser(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
