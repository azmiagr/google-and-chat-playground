package repository

import (
	"google-login/entity"
	"google-login/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByGoogleID(googleID string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	CreateUserFromOAuth(tx *gorm.DB, user *entity.User) (*entity.User, error)
	UpdateFromOAuth(tx *gorm.DB, user *entity.User) (*entity.User, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByGoogleID(googleID string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("google_id = ?", googleID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateUserFromOAuth(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateFromOAuth(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (*entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
