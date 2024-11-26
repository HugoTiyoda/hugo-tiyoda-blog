package authorsession

import (
	"blog/application/domain"
	"blog/application/ports"
	"fmt"
	"log"
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

// AuthorSessionService
func (service *AuthorSessionService) ValidateSession(token string) (*domain.AuthorSession, error) {
	// Buscar sessão pelo token
	session, err := service.authorSessionRepository.FindByToken(token)
	if err != nil {
		return nil, fmt.Errorf("session not found")
	}

	// Verificar se a sessão expirou
	if time.Now().After(session.ExpiresAt) {
		// Deletar sessão expirada
		if err := service.authorSessionRepository.Delete(session.Id); err != nil {
			// Loggar erro mas não retornar para o cliente
			log.Printf("failed to delete expired session: %v", err)
		}
		return nil, fmt.Errorf("session expired")
	}

	// Validar JWT
	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return session, nil
}

// Job para limpar sessões expiradas (pode rodar em uma goroutine)
func (service *AuthorSessionService) CleanExpiredSessions() error {
	return service.authorSessionRepository.DeleteAllExpired()
}
