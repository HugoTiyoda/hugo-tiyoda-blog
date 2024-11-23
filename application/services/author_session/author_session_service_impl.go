package authorsession

import (
	"blog/application/domain"
	"blog/application/ports"
	"time"

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

func (service *AuthorSessionService) Create(authorId string) (*domain.AuthorSession, error) {
	//TODO generate token JWT e user agent
	session := domain.AuthorSession{
		Id:        uuid.New().String(),
		AuthorId:  authorId,
		Token:     "",
		UserAgent: "",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}

	if err := service.authorSessionRepository.Save(&session); err != nil {
		return nil, err
	}

	return &session, nil
}
