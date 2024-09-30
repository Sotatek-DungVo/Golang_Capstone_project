package repository

import (
	"capstone_project/internal/models"
	"gorm.io/gorm"
)

type GameCategoryRepository struct {
	db *gorm.DB
}

func NewGameCategoryRepository(db *gorm.DB) *GameCategoryRepository {
	return &GameCategoryRepository{db: db}
}

func (r *GameCategoryRepository) List(page int, limit int) ([]*models.GameCategory, error) {
	var categories []*models.GameCategory
	offset := (page - 1) * limit

	if err := r.db.Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GameCategoryRepository) Create(category *models.GameCategory) error {
	return r.db.Create(category).Error
}
