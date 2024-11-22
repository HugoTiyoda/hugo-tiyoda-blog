package ports

import "blog/application/domain"

type BlogRepository interface {
	Save(d *domain.Post) error
	FindById(id string) (*domain.Post, error)
	FindByAuthorId(id string) ([]*domain.Post, error)
}
