package main

import (
	"github.com/emaanmohamed/chat-app/configs"
	"github.com/emaanmohamed/chat-app/internal/db"
	RoutesAuth "github.com/emaanmohamed/chat-app/routes/auth"
	RoutesMedia "github.com/emaanmohamed/chat-app/routes/media"
	RoutesMessage "github.com/emaanmohamed/chat-app/routes/message"
	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectToDB()
	server := gin.Default()

	server.Static("/static", "./web")
	server.Static("/uploads", "./web/uploads")

	routes := server.Group("v1/api")
	RoutesAuth.SetUpAuthRoutes(routes)
	RoutesMessage.SetUpMessageRoutes(routes)
	RoutesMedia.SetUpMediaRoutes(routes)
	server.Run(":" + configs.Envs.Port)

}
