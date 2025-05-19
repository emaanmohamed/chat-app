package main

import (
	"github.com/emaanmohamed/chat-app/configs"
	"github.com/emaanmohamed/chat-app/internal/db"
	RoutesAuth "github.com/emaanmohamed/chat-app/routes/auth"
	RoutesMedia "github.com/emaanmohamed/chat-app/routes/media"
	RoutesMessage "github.com/emaanmohamed/chat-app/routes/message"
	"github.com/gin-gonic/gin"
)

// @title Chat App API
// @version 1.0
// @description This is a chat application API.
// @termsOfService http://swagger.io/terms/
// @host localhost:8083
// @BasePath /v1/api

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
