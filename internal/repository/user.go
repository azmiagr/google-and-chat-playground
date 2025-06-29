package repository

import (
	"google-login/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByGoogleID(googleID string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	CreateUserFromOAuth(user *entity.User) (*entity.User, error)
	UpdateFromOAuth(user *entity.User) (*entity.User, error)
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

func (r *UserRepository) CreateUserFromOAuth(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateFromOAuth(user *entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
