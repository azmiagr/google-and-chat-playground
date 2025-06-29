package entity

import "time"

type User struct {
	UserID    int       `json:"id" gorm:"primaryKey"`
	GoogleID  *string   `json:"google_id" gorm:"uniqueIndex"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Picture   *string   `json:"picture"`
	Password  *string   `json:"-" gorm:"type:varchar(255)"`
	RoleID    uint      `json:"role_id" gorm:"not null;default:2"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Participants []Participant `json:"participants" gorm:"foreignKey:UserID"`
	Messages     []Message     `json:"messages" gorm:"foreignKey:UserID"`
}
