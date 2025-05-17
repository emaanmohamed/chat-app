package message

import (
	"github.com/emaanmohamed/chat-app/internal/db"
	"github.com/emaanmohamed/chat-app/internal/models"
	"github.com/emaanmohamed/chat-app/utils"
	"sync"
	"time"
)

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (ms *MessageService) SendMessage(request utils.MessageCreateRequest) (*models.Message, error) {
	message := models.Message{
		SenderID:   request.SenderID,
		ReceiverID: request.ReceiverID,
		Content:    request.Content,
		MediaURL:   request.MediaURL,
	}

	if err := db.DB.Create(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil

}

func (ms *MessageService) BroadcastMessage(request utils.BroadcastMessageRequest) ([]models.Message, error) {
	var users []models.User
	db.DB.Where("NOT id = ?", request.SenderID).Find(&users)

	var wg sync.WaitGroup
	var mu sync.Mutex
	messages := make([]models.Message, 0, len(users))
	errsChan := make(chan error, len(users))

	for _, user := range users {
		wg.Add(1)
		go func(user models.User) {
			defer wg.Done()
			msg := models.Message{
				SenderID:   request.SenderID,
				ReceiverID: user.ID,
				Content:    request.Content,
				MediaURL:   request.MediaURL,
				CreatedAt:  time.Now(),
			}
			if err := db.DB.Create(&msg).Error; err != nil {
				errsChan <- err
			}
			mu.Lock()
			messages = append(messages, msg)
			mu.Unlock()
		}(user)
	}

	wg.Wait()
	close(errsChan)

	if len(errsChan) > 0 {
		return messages, <-errsChan
	}
	return messages, nil
}
