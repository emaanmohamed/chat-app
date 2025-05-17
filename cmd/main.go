package main

import (
	"github.com/emaanmohamed/chat-app/configs"
	"github.com/emaanmohamed/chat-app/internal/db"
	RoutesAuth "github.com/emaanmohamed/chat-app/routes/auth"
	RoutesMessage "github.com/emaanmohamed/chat-app/routes/message"
	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectToDB()
	server := gin.Default()
	routes := server.Group("v1/api")
	RoutesAuth.SetUpAuthRoutes(routes)
	RoutesMessage.SetUpMessageRoutes(routes)
	server.Run(":" + configs.Envs.Port)

}
