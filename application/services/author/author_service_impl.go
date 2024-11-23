package author

import (
	"blog/application/domain"
	"blog/application/ports"
	authorsession "blog/application/services/author_session"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthorService struct {
	authorRepository ports.AuthorRepository
	sessionService   *authorsession.AuthorSessionService
}

func NewAuthorService(authorRepository ports.AuthorRepository, sessionService *authorsession.AuthorSessionService) *AuthorService {
	return &AuthorService{
		authorRepository: authorRepository,
		sessionService:   sessionService,
	}
}

func (s *AuthorService) Register(author *domain.Author, password string) (*domain.Author, string, error) {
	if exists, _ := s.authorRepository.ExistsByEmail(author.Email); exists {
		return nil, "", errors.New("email já cadastrado")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}
	author.Id = uuid.New().String()
	author.PasswordHash = string(hashedPassword)
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()
	author.IsActive = true

	if err := s.authorRepository.Save(author); err != nil {
		return nil, "", err
	}

	session, err := s.sessionService.Create(author.Id)
	if err != nil {
		return nil, "", err
	}

	return author, session.Token, nil
}
