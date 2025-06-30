package repository

import (
	"google-login/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IChatRepository interface {
	CreateMessage(tx *gorm.DB, message *entity.Message) error
	CreateConversation(tx *gorm.DB, convo *entity.Conversation) error
	CreateParticipant(tx *gorm.DB, participant *entity.Participant) error
	GetMessagesByConversationID(convoID uuid.UUID) ([]*entity.Message, error)
	GetConversationsByUser(userID int) ([]*entity.Conversation, error)
}

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) IChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) CreateMessage(tx *gorm.DB, message *entity.Message) error {
	err := r.db.Debug().Create(&message).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) CreateConversation(tx *gorm.DB, convo *entity.Conversation) error {
	err := r.db.Debug().Create(&convo).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) CreateParticipant(tx *gorm.DB, participant *entity.Participant) error {
	err := r.db.Debug().Create(&participant).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) GetMessagesByConversationID(convoID uuid.UUID) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.Debug().Preload("Conversation").
		Where("conversation_id = ?", convoID).
		Order("created_at asc").
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *ChatRepository) GetConversationsByUser(userID int) ([]*entity.Conversation, error) {
	var convos []*entity.Conversation
	err := r.db.Debug().
		Joins("JOIN participants on conversations.conversation_id = participants.conversation_id").
		Where("participants.user_id = ?", userID).
		Find(&convos).Error

	if err != nil {
		return nil, err
	}

	return convos, nil
}
