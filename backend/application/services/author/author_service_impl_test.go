package author

import (
	"blog/application/domain"
	"blog/application/ports/mocks"
	authorsession "blog/application/services/author_session"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	mockRepo := new(mocks.AuthorRepository)
	mockSessionRepo := new(mocks.AuthorSessionRepository)
	sessionService := authorsession.NewAuthorSessionService(mockSessionRepo)
	service := NewAuthorService(mockRepo, sessionService)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	author := &domain.Author{
		Id:           "123",
		Email:        "test@test.com",
		PasswordHash: string(hashedPassword),
		IsActive:     true,
	}

	mockRepo.On("FindByEmail", "test@test.com").Return(author, nil)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Author")).Return(nil)
	mockSessionRepo.On("Save", mock.AnythingOfType("*domain.AuthorSession")).Return(nil)

	result, err := service.Login("test@test.com", "password123", "test-agent", "127.0.0.1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, author.Id, result.Author.Id)
	mockRepo.AssertExpectations(t)
	mockSessionRepo.AssertExpectations(t)
}

func TestRegister(t *testing.T) {
	mockRepo := new(mocks.AuthorRepository)
	mockSessionRepo := new(mocks.AuthorSessionRepository)
	sessionService := authorsession.NewAuthorSessionService(mockSessionRepo)
	service := NewAuthorService(mockRepo, sessionService)

	newAuthor := &domain.Author{
		Email: "new@test.com",
		Name:  "Test User",
	}

	mockRepo.On("ExistsByEmail", "new@test.com").Return(false, nil)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Author")).Return(nil)
	mockSessionRepo.On("Save", mock.AnythingOfType("*domain.AuthorSession")).Return(nil)

	author, token, err := service.Register(newAuthor, "password123")

	assert.NoError(t, err)
	assert.NotNil(t, author)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, author.Id)
	assert.NotEmpty(t, author.PasswordHash)
	assert.True(t, author.IsActive)
	mockRepo.AssertExpectations(t)
	mockSessionRepo.AssertExpectations(t)
}
