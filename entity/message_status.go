package entity

import (
	"time"

	"github.com/google/uuid"
)

type MessageStatus struct {
	MessageStatusID uuid.UUID  `json:"message_status_id" gorm:"type:varchar(36);primaryKey"`
	MessageID       uuid.UUID  `json:"message_id"`
	UserID          uuid.UUID  `json:"user_id"`
	IsDelivered     bool       `json:"is_delivered" gorm:"type:bool"`
	IsSeen          bool       `json:"is_seen" gorm:"type:bool"`
	DeliveredAt     *time.Time `json:"delivered_at"`
	SeenAt          *time.Time `json:"seen_at"`

	Message Message `json:"message"`
	User    User    `json:"user"`
}
