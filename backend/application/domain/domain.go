package domain

import "time"

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Author struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Bio          string     `json:"bio"`
	PasswordHash string     `json:"-"`
	IsActive     bool       `json:"is_active"`
	LastLogin    *time.Time `json:"last_login"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type AuthorSession struct {
	Id        string    `json:"id"`
	AuthorId  string    `json:"author_id"`
	Token     string    `json:"token"`
	UserAgent string    `json:"user_agent"`
	IpAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type AuthorLogin struct {
	Author   Author    `json:"author"`
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type Comment struct {
	Id        string    `json:"id"`
	PostId    string    `json:"post_id"`
	AuthorId  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
