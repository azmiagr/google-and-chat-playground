package entity

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ConversationID uuid.UUID `json:"conversation_id" gorm:"type:varchar(36);primaryKey"`
	Title          string    `json:"title" gorm:"type:varchar(100)"`
	IsGroup        bool      `json:"is_group" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`

	Participants []Participant `json:"participants" gorm:"foreignKey:ConversationID"`
	Messages     []Message     `json:"messages" gorm:"foreignKey:ConversationID"`
}
