package postgresql

import (
	"blog/application/domain"
	"time"

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

func (r *AuthorSessionRepositoryImpl) DeleteAllExpired() error {
	query := `DELETE FROM author_sessions WHERE expires_at < $1`
	err := r.db.Exec(query, time.Now()).Error
	return err
}

func (r *AuthorSessionRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM author_sessions WHERE id = $1`
	err := r.db.Exec(query, id).Error
	return err
}

func (r *AuthorSessionRepositoryImpl) FindByToken(token string) (*domain.AuthorSession, error) {
	var e AuthorSession
	err := r.db.First(&e, "token = ?", token).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return e.ToDomain(), err
}
