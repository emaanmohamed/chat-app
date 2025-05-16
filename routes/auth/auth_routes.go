package auth

import (
	"github.com/emaanmohamed/chat-app/controllers"
	"github.com/emaanmohamed/chat-app/internal/services/auth"
	"github.com/gin-gonic/gin"
)

var (
	authService    = auth.NewAuthService()
	authController = controllers.NewAuthController(authService)
)

func SetUpAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}

}
