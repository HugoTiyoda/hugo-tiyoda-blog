package mocks

import (
	"blog/application/domain"

	"github.com/stretchr/testify/mock"
)

type BlogRepository struct {
	mock.Mock
}

func (m *BlogRepository) Save(post *domain.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *BlogRepository) FindById(id string) (*domain.Post, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Post), args.Error(1)
}

func (m *BlogRepository) FindByAuthorId(authorId string) ([]*domain.Post, error) {
	args := m.Called(authorId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Post), args.Error(1)
}
