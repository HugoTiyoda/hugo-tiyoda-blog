package postgresql

import (
	"blog/application/domain"
	"log"

	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepositoryImpl(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) Save(d *domain.Post) error {
	entity := NewPost(d)
	log.Println("[INFO] SALVO")
	return r.db.Save(entity).Error
}

func (r *PostRepositoryImpl) FindById(id string) (*domain.Post, error) {
	var e Post
	err := r.db.First(&e, id).Error
	if err != nil {
		return nil, err
	}

	return e.ToDomain(), err
}

func (r *PostRepositoryImpl) FindByAuthorId(id string) ([]*domain.Post, error) {
	var entities []Post
	query := "select * from posts where author_id = ? "
	if err := r.db.Raw(query, id).Scan(&entities).Error; err != nil {
		return nil, err
	}

	domains := make([]*domain.Post, len(entities))
	for _, e := range entities {
		domains = append(domains, e.ToDomain())
	}
	return domains, nil
}
