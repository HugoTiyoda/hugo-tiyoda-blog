package postgresql

import (
	"blog/application/domain"
	"time"
)

type Post struct {
	Id        string `gorm:"primaryKey"`
	Title     string
	Content   string
	AuthorId  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(p *domain.Post) *Post {
	return &Post{
		Id:        p.Id,
		Title:     p.Title,
		Content:   p.Content,
		AuthorId:  p.AuthorId,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (p *Post) ToDomain() *domain.Post {
	return &domain.Post{
		Id:        p.Id,
		Title:     p.Title,
		Content:   p.Content,
		AuthorId:  p.AuthorId,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

type Author struct {
	Id           string `gorm:"primaryKey"`
	Name         string
	Email        string
	Bio          string
	PasswordHash string
	IsActive     bool
	LastLogin    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewAuthor(a *domain.Author) *Author {
	return &Author{
		Id:           a.Id,
		Name:         a.Name,
		Email:        a.Email,
		Bio:          a.Bio,
		PasswordHash: a.PasswordHash,
		IsActive:     a.IsActive,
		LastLogin:    a.LastLogin,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
}

func (a *Author) ToDomain() *domain.Author {
	return &domain.Author{
		Id:           a.Id,
		Name:         a.Name,
		Email:        a.Email,
		Bio:          a.Bio,
		PasswordHash: a.PasswordHash,
		IsActive:     a.IsActive,
		LastLogin:    a.LastLogin,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
}

type AuthorSession struct {
	Id        string
	AuthorId  string
	Token     string
	UserAgent string
	IpAddress string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewAuthorSession(a *domain.AuthorSession) *AuthorSession {
	return &AuthorSession{
		Id:        a.Id,
		AuthorId:  a.AuthorId,
		Token:     a.Token,
		UserAgent: a.UserAgent,
		IpAddress: a.IpAddress,
		CreatedAt: a.CreatedAt,
		ExpiresAt: a.ExpiresAt,
	}
}
