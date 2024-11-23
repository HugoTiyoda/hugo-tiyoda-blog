package ports

import "blog/application/domain"

type AuthorService interface {
	Register(author *domain.Author, password string) (*domain.Author, string, error)
}
type PostService interface {
	Create(post *domain.Post) error
	Update(id, tittle, content string) error
	FindByAuthorId(id string) ([]*domain.Post, error)
}
