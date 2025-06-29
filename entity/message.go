package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageID      uuid.UUID `json:"message_id" gorm:"type:varchar(36);primaryKey"`
	ConversationID uuid.UUID `json:"conversation_id"`
	UserID         uuid.UUID `json:"user_id"`
	Content        string    `json:"content" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`

	Conversation Conversation    `json:"conversation"`
	User         User            `json:"user"`
	Status       []MessageStatus `gorm:"foreignKey:MessageID"`
}
