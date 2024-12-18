package mocks

import (
	"blog/application/domain"

	"github.com/stretchr/testify/mock"
)

type AuthorSessionRepository struct {
	mock.Mock
}

func (m *AuthorSessionRepository) Save(session *domain.AuthorSession) error {
	args := m.Called(session)
	return args.Error(0)
}

func (m *AuthorSessionRepository) DeleteAllExpired() error {
	args := m.Called()
	return args.Error(0)
}

func (m *AuthorSessionRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *AuthorSessionRepository) FindByToken(token string) (*domain.AuthorSession, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AuthorSession), args.Error(1)
}
