package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"google-login/entity"
	"google-login/internal/repository"
	"google-login/model"
	"google-login/pkg/bcrypt"
	"google-login/pkg/config"
	"google-login/pkg/database/mariadb"
	"google-login/pkg/jwt"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type IOAuthService interface {
	GetGoogleLoginURL() (string, string, error)
	HandleGoogleCallback(code, state, savedState string) (*model.OAuthLoginResponse, error)
	GetUserInfoFromGoogle(accessToken string) (*model.GoogleUserInfo, error)
}

type OAuthService struct {
	db             *gorm.DB
	UserRepository repository.IUserRepository
	oauth          *config.OAuthConfig
	bcrypt         bcrypt.Interface
	jwtAuth        jwt.Interface
}

func NewOAuthService(userRepo repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface, oauth *config.OAuthConfig) IOAuthService {
	return &OAuthService{
		db:             mariadb.Connection,
		UserRepository: userRepo,
		bcrypt:         bcrypt,
		jwtAuth:        jwt,
		oauth:          oauth,
	}
}

func (s *OAuthService) GetGoogleLoginURL() (string, string, error) {
	state, err := s.generateState()
	if err != nil {
		return "", "", err
	}

	url := s.oauth.GoogleConfig.AuthCodeURL(state)
	return url, state, nil
}

func (s *OAuthService) HandleGoogleCallback(code, state, savedState string) (*model.OAuthLoginResponse, error) {
	if state != savedState {
		return nil, errors.New("invalid state parameter")
	}

	token, err := s.oauth.GoogleConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.New("failed to exchange code for token")
	}

	googleUser, err := s.GetUserInfoFromGoogle(token.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.findOrCreateUser(googleUser)
	if err != nil {
		return nil, err
	}

	jwtToken, err := s.jwtAuth.CreateToken(user.UserID)
	if err != nil {
		return nil, errors.New("failed to create JWT token")
	}

	return &model.OAuthLoginResponse{
		Token: jwtToken,
		User:  user,
	}, nil
}

func (s *OAuthService) GetUserInfoFromGoogle(accessToken string) (*model.GoogleUserInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get user info from Google")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo model.GoogleUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func (s *OAuthService) findOrCreateUser(googleUser *model.GoogleUserInfo) (*entity.User, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	user, err := s.UserRepository.FindByGoogleID(googleUser.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user != nil {
		user.Name = googleUser.Name
		user.Picture = &googleUser.Picture
		_, err = s.UserRepository.UpdateFromOAuth(tx, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	existingUser, err := s.UserRepository.FindByEmail(googleUser.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser != nil {
		existingUser.GoogleID = &googleUser.ID
		existingUser.Picture = &googleUser.Picture
		_, err = s.UserRepository.UpdateFromOAuth(tx, existingUser)
		if err != nil {
			return nil, err
		}
		return existingUser, nil
	}

	newUser := &entity.User{
		GoogleID: &googleUser.ID,
		Email:    googleUser.Email,
		Name:     googleUser.Name,
		Picture:  &googleUser.Picture,
		RoleID:   2,
	}

	_, err = s.UserRepository.CreateUserFromOAuth(tx, newUser)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return s.UserRepository.FindByGoogleID(googleUser.ID)
}

func (s *OAuthService) generateState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
