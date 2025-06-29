package jwt

import (
	"errors"
	"google-login/entity"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Interface interface {
	CreateToken(userID int) (string, error)
	ValidateToken(tokenString string) (int, error)
	GetLoginUSer(ctx *gin.Context) (entity.User, error)
}

type jsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

func Init() Interface {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatalf("error init jwt %v", err)
	}

	return &jsonWebToken{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expTime) * time.Hour,
	}
}

func (j *jsonWebToken) CreateToken(userID int) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func (j *jsonWebToken) ValidateToken(tokenString string) (int, error) {
	var (
		claim  Claims
		userID int
	)

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return userID, err
	}

	if !token.Valid {
		return userID, errors.New("token is not valid")
	}

	userID = claim.UserID

	return userID, nil
}

func (j *jsonWebToken) GetLoginUSer(ctx *gin.Context) (entity.User, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return entity.User{}, errors.New("failed to get login user")
	}

	return user.(entity.User), nil
}
