package entity

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	ParticipantID  uuid.UUID `json:"participant_id" gorm:"type:varchar(36);primaryKey"`
	UserID         int       `json:"user_id"`
	ConversationID uuid.UUID `json:"conversation_id"`
	IsAdmin        bool      `json:"is_admin" gorm:"type:bool"`
	JoinedAt       time.Time `json:"joined_at"`

	User         User         `json:"user"`
	Conversation Conversation `json:"conversation"`
}
