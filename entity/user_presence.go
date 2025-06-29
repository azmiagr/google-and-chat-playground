package entity

import "time"

type UserPresence struct {
	UserID   int       `json:"user_id" gorm:"type:int;primaryKey"`
	IsBool   bool      `json:"is_bool" gorm:"type:bool"`
	LastSeen time.Time `json:"last_seen" gorm:"type:timestamp"`

	User User `json:"user"`
}
