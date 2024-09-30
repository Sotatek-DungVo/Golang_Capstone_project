package repository

import (
	"capstone_project/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetDB() *gorm.DB {
	return r.db
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetUserByEmailOrUsername(identifier string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	return &user, err
}

func (r *UserRepository) List(page int, limit int, search map[string]string) ([]*models.User, error) {
	var users []*models.User
	offset := (page - 1) * limit

	query := r.db.Model(&models.User{})

	if username, ok := search["username"]; ok && username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if gender, ok := search["gender"]; ok && gender != "" {
		query = query.Where("gender = ?", models.Gender(gender))
	}

	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
