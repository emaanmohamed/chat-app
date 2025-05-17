package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SenderID   uuid.UUID `gorm:"not null"`
	ReceiverID uuid.UUID
	Content    string
	MediaURL   *string
	CreatedAt  time.Time
}
