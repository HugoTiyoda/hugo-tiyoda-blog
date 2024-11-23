package postgresql

import (
	"blog/application/domain"

	"gorm.io/gorm"
)

type AuthorSessionRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthorSessionRepositoryImpl(db *gorm.DB) *AuthorSessionRepositoryImpl {
	return &AuthorSessionRepositoryImpl{db: db}
}

func (r *AuthorSessionRepositoryImpl) Save(d *domain.AuthorSession) error {
	entity := NewAuthorSession(d)
	return r.db.Save(entity).Error
}

