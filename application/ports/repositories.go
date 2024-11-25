package ports

import "blog/application/domain"

type AuthorRepository interface {
	Save(d *domain.Author) error
	FindById(id string) (*domain.Author, error)
	ExistsByEmail(email string) (bool, error)
	FindByEmail(email string) (*domain.Author, error)
}

type AuthorSessionRepository interface {
	Save(d *domain.AuthorSession) error
}

type BlogRepository interface {
	Save(d *domain.Post) error
	FindById(id string) (*domain.Post, error)
	FindByAuthorId(id string) ([]*domain.Post, error)
}
