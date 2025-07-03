package service

import (
	"errors"
	"google-login/entity"
	"google-login/internal/repository"
	"google-login/model"
	"google-login/pkg/bcrypt"
	"google-login/pkg/database/mariadb"
	"google-login/pkg/jwt"

	"gorm.io/gorm"
)

type IUserService interface {
	Register(param model.UserRegisterParam) error
	Login(param model.UserLoginParam) (*model.LoginResponse, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserService struct {
	db             *gorm.DB
	bcrypt         bcrypt.Interface
	jwt            jwt.Interface
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface) IUserService {
	return &UserService{
		db:             mariadb.Connection,
		jwt:            jwt,
		bcrypt:         bcrypt,
		UserRepository: userRepository,
	}
}

func (s *UserService) Register(param model.UserRegisterParam) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	existingUser, err := s.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existingUser != nil {
		return errors.New("email already exists")
	}

	if param.Password != param.ConfirmPassword {
		return errors.New("password doesn't match")
	}

	hash, err := s.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: &hash,
		RoleID:   2,
	}

	_, err = s.UserRepository.CreateUserFromOAuth(tx, user)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(param model.UserLoginParam) (*model.LoginResponse, error) {
	user, err := s.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = s.bcrypt.CompareAndHashPassword(*user.Password, *param.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwt.CreateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	result := &model.LoginResponse{
		Token: token,
	}

	return result, nil
}

func (s *UserService) GetUser(param model.UserParam) (*entity.User, error) {
	return s.UserRepository.GetUser(param)
}
