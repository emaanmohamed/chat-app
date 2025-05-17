package message

import (
	"github.com/emaanmohamed/chat-app/controllers"
	"github.com/emaanmohamed/chat-app/internal/services/message"
	"github.com/emaanmohamed/chat-app/middleware"
	"github.com/gin-gonic/gin"
)

var (
	messageService    = message.NewMessageService()
	messageController = controllers.NewMessageController(messageService)
)

func SetUpMessageRoutes(router *gin.RouterGroup) {
	messageGroup := router.Group("/message").Use(middleware.AuthMiddleware())
	{
		messageGroup.POST("/send", messageController.SendMessage)
		messageGroup.POST("/broadcast", messageController.BroadcastMessage)

	}

}
