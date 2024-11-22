package ports

import "blog/application/domain"

type PostService interface {
	Create(post *domain.Post) error
	Update(id, tittle, content string) error
	FindByAuthorId(id string) ([]*domain.Post, error)
}
