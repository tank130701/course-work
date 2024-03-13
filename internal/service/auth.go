package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository"
	"time"
)

const (
	salt       = "b2ee2620a71de50ef9731232ae7e860a240f6859b558c279035fe28e26cfed1f"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.IAuthorization
}

func NewAuthService(repo repository.IAuthorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// func (s *AuthService) GenerateToken(username, password string) (string, error) {
// 	user, error := s.repo.GetUser(username, generatePasswordHash(password))
// 	if error != nil {
// 		return "", error
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		},
// 		user.Id,
// 	})

// 	return token.SignedString([]byte(signingKey))
// }

func (s *AuthService) GenerateToken(username, password string) (string, string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", "", err
	}

	// Создание access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	accessString, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	// Создание refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)), // Refresh токен действителен в течение недели
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	refreshString, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	// Разбор refresh токена
	token, err := jwt.ParseWithClaims(refreshToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	// Создание нового access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		claims.UserId,
	})

	accessString, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return accessString, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
