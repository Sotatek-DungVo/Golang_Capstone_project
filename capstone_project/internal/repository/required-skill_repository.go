package repository

import (
	"capstone_project/internal/models"
	"gorm.io/gorm"
)

type RequiredSkillRepository struct {
	db *gorm.DB
}

func NewRequiredSkillRepository(db *gorm.DB) *RequiredSkillRepository {
	return &RequiredSkillRepository{db: db}
}

func (r *RequiredSkillRepository) Create(skill *models.RequiredSkill) error {
	return r.db.Create(skill).Error
}

func (r *RequiredSkillRepository) FindAll() ([]models.RequiredSkill, error) {
	var skills []models.RequiredSkill
	err := r.db.Find(&skills).Error
	return skills, err
}
