package controllers

import (
	"github.com/emaanmohamed/chat-app/internal/services/message"
	"github.com/emaanmohamed/chat-app/utils"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService *message.MessageService
}

func NewMessageController(messageService *message.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

func (messageController *MessageController) SendMessage(c *gin.Context) {
	var sendMessageRequest utils.MessageCreateRequest
	if err := c.ShouldBindJSON(&sendMessageRequest); err != nil {
		utils.RespondWithError(c, 400, "Invalid request")
		return
	}
	msg, err := messageController.messageService.SendMessage(sendMessageRequest)
	if err != nil {
		utils.RespondWithError(c, 500, err.Error())
		return
	}
	utils.RespondWithJSON(c, 200, msg)
}

func (messageController *MessageController) BroadcastMessage(c *gin.Context) {
	var broadcastMessageRequest utils.BroadcastMessageRequest
	if err := c.ShouldBindJSON(&broadcastMessageRequest); err != nil {
		utils.RespondWithError(c, 400, "Invalid request")
		return
	}
	msg, err := messageController.messageService.BroadcastMessage(broadcastMessageRequest)
	if err != nil {
		utils.RespondWithError(c, 500, err.Error())
		return
	}
	utils.RespondWithJSON(c, 200, msg)

}
