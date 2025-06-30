package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	ChatRepository IChatRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		ChatRepository: NewChatRepository(db),
	}
}
