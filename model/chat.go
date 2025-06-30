package model

import "github.com/google/uuid"

type SendMessageInput struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	UserID         int       `json:"user_id"`
	Content        string    `json:"content"`
}

type CreateConversationInput struct {
	UserIDs []int  `json:"user_ids"`
	Title   string `json:"title"`
}
