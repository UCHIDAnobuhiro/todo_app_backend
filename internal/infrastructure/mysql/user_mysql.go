package mysql

import (
	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"

	"gorm.io/gorm"
)

type userMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) repository.UserRepository {
	return &userMySQL{db: db}
}

func (r *userMySQL) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

func (r *userMySQL) FindByEmail(email string) (*domain.User, error) {
	var u domain.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userMySQL) FindByID(id uint) (*domain.User, error) {
	var u domain.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
