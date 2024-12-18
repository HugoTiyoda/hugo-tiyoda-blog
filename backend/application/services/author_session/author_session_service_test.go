package authorsession

import (
	"blog/application/ports/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSession(t *testing.T) {
	mockRepo := new(mocks.AuthorSessionRepository)
	service := NewAuthorSessionService(mockRepo)

	mockRepo.On("Save", mock.AnythingOfType("*domain.AuthorSession")).Return(nil)

	session, err := service.Create("author123", "test-agent", "127.0.0.1")

	assert.NoError(t, err)
	assert.NotNil(t, session)
	assert.NotEmpty(t, session.Id)
	assert.NotEmpty(t, session.Token)
	assert.Equal(t, "author123", session.AuthorId)
	assert.Equal(t, "test-agent", session.UserAgent)
	assert.Equal(t, "127.0.0.1", session.IpAddress)
	assert.True(t, session.ExpiresAt.After(time.Now()))
	mockRepo.AssertExpectations(t)
}

func TestValidateSession(t *testing.T) {
    mockRepo := new(mocks.AuthorSessionRepository)
    service := NewAuthorSessionService(mockRepo)

    mockRepo.On("Save", mock.AnythingOfType("*domain.AuthorSession")).Return(nil)

    validSession, err := service.Create("author123", "test-agent", "127.0.0.1")
    if err != nil {
		t.Fatalf("Failed to create valid session: %v", err)
	}
    mockRepo.On("FindByToken", validSession.Token).Return(validSession, nil)

    session, err := service.ValidateSession(validSession.Token)

    assert.NoError(t, err)
    assert.NotNil(t, session)
    assert.Equal(t, validSession.Id, session.Id)
    mockRepo.AssertExpectations(t)
}

func TestCleanExpiredSessions(t *testing.T) {
	mockRepo := new(mocks.AuthorSessionRepository)
	service := NewAuthorSessionService(mockRepo)

	mockRepo.On("DeleteAllExpired").Return(nil)

	err := service.CleanExpiredSessions()

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
