package service

import (
	"google-login/entity"
	"google-login/internal/repository"
	"google-login/model"
	"google-login/pkg/database/mariadb"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IChatService interface {
	GetMessagesByConversationID(convoID uuid.UUID) ([]*entity.Message, error)
	SendMessage(param model.SendMessageInput, convoID uuid.UUID) error
	CreateConversation(param model.CreateConversationInput) error
	GetConversationsByUser(userID int) ([]*entity.Conversation, error)
}

type ChatService struct {
	db             *gorm.DB
	ChatRepository repository.IChatRepository
}

func NewChatService(chatRepository repository.IChatRepository) IChatService {
	return &ChatService{
		db:             mariadb.Connection,
		ChatRepository: chatRepository,
	}
}

func (s *ChatService) CreateConversation(param model.CreateConversationInput) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	convo := &entity.Conversation{
		ConversationID: uuid.New(),
		Title:          param.Title,
		IsGroup:        len(param.UserIDs) > 2,
		CreatedAt:      time.Now(),
	}

	err := s.ChatRepository.CreateConversation(tx, convo)
	if err != nil {
		return err
	}

	for _, userID := range param.UserIDs {
		participant := &entity.Participant{
			ParticipantID:  uuid.New(),
			ConversationID: convo.ConversationID,
			UserID:         userID,
			JoinedAt:       time.Now(),
		}
		err := s.ChatRepository.CreateParticipant(tx, participant)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ChatService) SendMessage(param model.SendMessageInput, convoID uuid.UUID) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	msg := &entity.Message{
		MessageID:      uuid.New(),
		ConversationID: convoID,
		UserID:         param.UserID,
		Content:        param.Content,
		CreatedAt:      time.Now(),
	}

	err := s.ChatRepository.CreateMessage(tx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *ChatService) GetMessagesByConversationID(convoID uuid.UUID) ([]*entity.Message, error) {
	convos, err := s.ChatRepository.GetMessagesByConversationID(convoID)
	if err != nil {
		return nil, err
	}

	return convos, nil
}

func (s *ChatService) GetConversationsByUser(userID int) ([]*entity.Conversation, error) {
	convos, err := s.ChatRepository.GetConversationsByUser(userID)
	if err != nil {
		return nil, err
	}

	return convos, nil
}
