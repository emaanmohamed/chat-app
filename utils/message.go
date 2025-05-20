package utils

import "github.com/google/uuid"

type MessageCreateRequest struct {
	ReceiverID uuid.UUID `json:"receiver_id" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	MediaURL   *string   `json:"media_url"`
}

type BroadcastMessageRequest struct {
	Content  string  `json:"content" binding:"required"`
	MediaURL *string `json:"media_url"`
}
