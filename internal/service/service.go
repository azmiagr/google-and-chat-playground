package service

import (
	"google-login/internal/repository"
	"google-login/pkg/bcrypt"
	"google-login/pkg/config"
	"google-login/pkg/jwt"
)

type Service struct {
	OAuthService IOAuthService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface, oauth *config.OAuthConfig) *Service {
	return &Service{
		OAuthService: NewOAuthService(repository.UserRepository, bcrypt, jwt, oauth),
	}
}
