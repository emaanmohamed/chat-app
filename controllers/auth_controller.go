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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body utils.AuthRequest true "User registration details"
// @Success 200 {object} utils.Response{data=utils.AuthRequest}
// @Failure 400 {object} utils.Response{error=string}
// @Router /auth/register [post]

func (authController AuthController) Register(c *gin.Context) {
	var registerRequest utils.AuthRequest
	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := authController.authService.Register(registerRequest)
	if err != nil {
		utils.RespondWithError(c, 400, "Error registering user")
		return
	}
	utils.RespondWithJSON(c, 200, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body utils.AuthRequest true "User login details"
// @Success 200 {object} utils.Response{data=utils.AuthRequest}
// @Failure 400 {object} utils.Response{error=string}
// @Router /auth/login [post]

func (authController *AuthController) Login(c *gin.Context) {
	var loginRequest utils.AuthRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.RespondWithError(c, 400, "Invalid request")
		return
	}

	token, user, err := authController.authService.Login(loginRequest)
	if err != nil {
		utils.RespondWithError(c, 400, "Invalid credentials")
		return
	}
	utils.RespondWithJSON(c, 200, gin.H{"token": token, "user_id": user.ID.String()})
}
