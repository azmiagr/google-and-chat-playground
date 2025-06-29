package service

import (
	"google-login/entity"
	"google-login/internal/repository"
	"google-login/model"
	"google-login/pkg/database/mariadb"

	"gorm.io/gorm"
)

type IUserService interface {
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserService struct {
	db             *gorm.DB
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		db:             mariadb.Connection,
		UserRepository: userRepository,
	}
}

func (u *UserService) GetUser(param model.UserParam) (*entity.User, error) {
	return u.UserRepository.GetUser(param)
}
