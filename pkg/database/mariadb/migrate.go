package mariadb

import (
	"google-login/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Conversation{},
		&entity.Participant{},
		&entity.Message{},
		&entity.MessageStatus{},
		&entity.UserPresence{},
	); err != nil {
		return err
	}
	return nil
}
