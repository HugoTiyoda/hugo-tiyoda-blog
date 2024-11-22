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
	Id    string
	Name  string
	Email string
	Bio   string
}
