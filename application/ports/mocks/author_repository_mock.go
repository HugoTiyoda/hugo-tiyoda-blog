package mocks

import (
	"blog/application/domain"

	"github.com/stretchr/testify/mock"
)

type AuthorRepository struct {
	mock.Mock
}

func (m *AuthorRepository) Save(author *domain.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *AuthorRepository) FindById(id string) (*domain.Author, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Author), args.Error(1)
}

func (m *AuthorRepository) ExistsByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

func (m *AuthorRepository) FindByEmail(email string) (*domain.Author, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Author), args.Error(1)
}
