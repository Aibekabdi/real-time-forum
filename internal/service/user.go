package service

import (
	"context"
	"errors"
	"fmt"
	"forum/internal/models"
	"forum/internal/repository"
	"forum/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	userRepo  repository.User
	secretKey string
}

func newAuthService(userRepo repository.User, secretKey string) *AuthService {
	return &AuthService{userRepo: userRepo, secretKey: secretKey}
}

func (s *AuthService) Create(ctx context.Context, user models.User) error {
	if err := utils.IsValidRegister(&user); err != nil {
		return fmt.Errorf("user service: sign in: %w", err)
	}
	return s.userRepo.Create(ctx, user)
}

func (s *AuthService) SignIn(ctx context.Context, user models.User) (string, error) {
	id, hash, err := s.userRepo.GetUser(ctx, user.Nickname)
	if err != nil {
		return "", fmt.Errorf("user service: sign in: %w", err)
	}

	if err = utils.CompareHashAndPassword(hash, user.Password); err != nil {
		return "", fmt.Errorf("user service: sign in: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserToken: models.UserToken{
			UserId: id,
			Role:   models.Roles.User,
		},
	})

	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthService) ParseToken(accessToken string) (models.UserToken, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return models.UserToken{}, err
	}
	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return models.UserToken{}, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserToken, nil
}
