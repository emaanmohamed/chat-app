package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Content    string    `json:"content"`
	MediaURL   *string   `json:"media_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
