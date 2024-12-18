package author

import (
	"blog/application/domain"
	"blog/application/ports"
	authorsession "blog/application/services/author_session"
	"errors"
	"fmt"
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
func (service *AuthorService) Login(email, password, userAgent, ipAddress string) (*domain.AuthorLogin, error) {
	author, err := service.authorRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !author.IsActive {
		return nil, errors.New("account is disabled")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(author.PasswordHash),
		[]byte(password),
	); err != nil {
		return nil, errors.New("invalid credentials")
	}

	session, err := service.sessionService.Create(author.Id, userAgent, ipAddress)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	author.LastLogin = &now
	if err := service.authorRepository.Save(author); err != nil {
		return nil, fmt.Errorf("failed to update last login: %v", err)
	}

	return &domain.AuthorLogin{
		Author:   *author,
		Token:    session.Token,
		ExpireAt: session.ExpiresAt,
	}, nil
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

	//Na hora que registra já cria a sessão pra continuidade do fluxo. Ex.: pagina de login -> tela inicial
	session, err := s.sessionService.Create(author.Id, "", "")
	if err != nil {
		return nil, "", err
	}

	return author, session.Token, nil
}
