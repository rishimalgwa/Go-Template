package users

import (
	"github.com/google/uuid"
	"github.com/rishimalgwa/go-template/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) Find(id *uuid.UUID) (*models.Users, error) {
	u := &models.Users{}
	result := r.DB.Where("id = ?", id).First(u)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
