package authorsession

import (
	"blog/application/domain"
	"blog/application/ports"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthorSessionService struct {
	authorSessionRepository ports.AuthorSessionRepository
}

func NewAuthorSessionService(authorSessionRepository ports.AuthorSessionRepository) *AuthorSessionService {
	return &AuthorSessionService{
		authorSessionRepository: authorSessionRepository,
	}
}

func (service *AuthorSessionService) Create(authorId, userAgent, ipAddress string) (*domain.AuthorSession, error) {
	tokenClaims := jwt.MapClaims{
		"sub": authorId,                                // subject (autor)
		"iat": time.Now().Unix(),                       // issued at
		"exp": time.Now().Add(time.Minute * 30).Unix(), // expiration
		"sid": uuid.New().String(),                     // session id
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}
	session := domain.AuthorSession{
		Id:        uuid.New().String(),
		AuthorId:  authorId,
		Token:     tokenString,
		UserAgent: userAgent,
		IpAddress: ipAddress,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}

	if err := service.authorSessionRepository.Save(&session); err != nil {
		return nil, err
	}

	return &session, nil
}
