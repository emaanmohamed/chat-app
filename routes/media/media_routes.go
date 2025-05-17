package media

import (
	"github.com/emaanmohamed/chat-app/controllers"
	"github.com/emaanmohamed/chat-app/middleware"
	"github.com/gin-gonic/gin"
)

var (
	mediaController = controllers.NewMediaController("./media/uploads")
)

func SetUpMediaRoutes(router *gin.RouterGroup) {
	mediaGroup := router.Group("/media").Use(middleware.AuthMiddleware())
	{
		mediaGroup.POST("/upload", mediaController.UploadMediaHandler())

	}

}
