package postgresql

import (
	"blog/application/domain"

	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthorRepositoryImpl(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{db: db}
}

func (r *AuthorRepositoryImpl) Save(d *domain.Author) error {
	entity := NewAuthor(d)
	return r.db.Save(entity).Error
}

func (r *AuthorRepositoryImpl) FindById(id string) (*domain.Author, error) {
	var e Author
	err := r.db.First(&e, id).Error
	if err != nil {
		return nil, err
	}

	return e.ToDomain(), err
}

func (r *AuthorRepositoryImpl) ExistsByEmail(email string) (bool, error) {
	query := "select count(*) from authors where email = ?"
	var count int64
	if err := r.db.Raw(query, email).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *AuthorRepositoryImpl) FindByEmail(email string) (*domain.Author, error) {
	var e Author
	err := r.db.First(&e, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return e.ToDomain(), err
}
