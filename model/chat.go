package model

import (
	"time"

	"github.com/google/uuid"
)

type SendMessageInput struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

type CreateConversationInput struct {
	UserIDs []int  `json:"user_ids"`
	Title   string `json:"title"`
}

type GetMessageResponse struct {
	ConversationTitle string    `json:"title"`
	MessageID         uuid.UUID `json:"message_id"`
	SenderID          int       `json:"sender_id"`
	Content           string    `json:"content"`
	CreatedAt         time.Time `json:"created_at"`
}

type GetConversationResponse struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	Title          string    `json:"title"`
	IsGroup        bool      `json:"is_group"`
	CreatedAt      time.Time `json:"created_at"`
}
