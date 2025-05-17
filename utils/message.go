package utils

import "github.com/google/uuid"

type MessageCreateRequest struct {
	SenderID   uuid.UUID `json:"sender_id" binding:"required"`
	ReceiverID uuid.UUID `json:"receiver_id" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	MediaURL   *string   `json:"media_url"`
}

type BroadcastMessageRequest struct {
	SenderID uuid.UUID `json:"sender_id" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	MediaURL *string   `json:"media_url"`
}
