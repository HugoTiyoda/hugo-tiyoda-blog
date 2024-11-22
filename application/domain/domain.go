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
	Id    string
	Name  string
	Email string
	Bio   string
}

type Comment struct {
	Id        string    `json:"id"`
	PostId    string    `json:"post_id"`
	AuthorId  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
