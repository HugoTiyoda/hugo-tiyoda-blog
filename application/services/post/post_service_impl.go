package post

import (
	"blog/application/domain"
	"blog/application/ports"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	blogRepository ports.BlogRepository
}

func NewPostService(blogRepository ports.BlogRepository) *PostService {
	return &PostService{
		blogRepository: blogRepository,
	}
}
func (s *PostService) FindByAuthorId(id string) ([]*domain.Post, error) {
	return s.blogRepository.FindByAuthorId(id)
}

func (s *PostService) Create(post *domain.Post) error {
	if post.Id == "" {
		post.Id = uuid.New().String()
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if err := s.blogRepository.Save(post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) Update(id, tittle, content string) error {
	post, err := s.blogRepository.FindById(id)
	if err != nil {
		return err
	}

	if tittle != "" {
		post.Title = tittle
	}
	if content != "" {
		post.Content = content
	}

	post.UpdatedAt = time.Now()

	if err := s.blogRepository.Save(post); err != nil {
		return err
	}

	return nil
}
