package controllers

import (
	"github.com/emaanmohamed/chat-app/internal/services/auth"
	"github.com/emaanmohamed/chat-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *auth.AuthService
}

func NewAuthController(authService *auth.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (authController *AuthController) Register(c *gin.Context) {
	var request utils.AuthRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := authController.authService.Register(request)
	if err != nil {
		utils.RespondWithError(c, 400, "Error registering user")
		return
	}
	utils.RespondWithJSON(c, 200, gin.H{"message": "User registered successfully"})
}
