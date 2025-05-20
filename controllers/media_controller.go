package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

type MediaController struct {
	UploadPath string
}

func NewMediaController(uploadPath string) *MediaController {
	return &MediaController{UploadPath: uploadPath}

}
func (mediaController *MediaController) UploadMediaHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
			return
		}

		filename := uuid.New().String() + filepath.Ext(file.Filename)
		fullPath := filepath.Join(mediaController.UploadPath, filename)

		err = c.SaveUploadedFile(file, fullPath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}

		c.JSON(200, gin.H{"message": "File uploaded successfully", "filename": file.Filename, "url": "http://localhost:8080/media/uploads/" + filename})
	}

}
