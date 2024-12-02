package ports

import "blog/application/domain"

type AuthorService interface {
	Login(email, password, userAgent, ipAddress string) (*domain.AuthorLogin, error)
	Register(author *domain.Author, password string) (*domain.Author, string, error)
}
type PostService interface {
	Create(post *domain.Post) error
	Update(id, tittle, content string) error
	FindByAuthorId(id string) ([]*domain.Post, error)
}

type SessionService interface {
	Create(authorId, userAgent, ipAddress string) (*domain.AuthorSession, error)
	ValidateSession(token string) (*domain.AuthorSession, error)
	CleanExpiredSessions() error
}
