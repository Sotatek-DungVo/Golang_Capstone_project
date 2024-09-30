package repository

import (
	"capstone_project/internal/models"
	"gorm.io/gorm"
)

type GameRequestRepository interface {
	Create(gameRequest *models.GameRequest) (*models.GameRequest, error)
	GetByID(id uint) (*models.GameRequest, error)
	Update(gameRequest *models.GameRequest) (*models.GameRequest, error)
	GetByUserAndGame(userID uint, gameID uint) (*models.GameRequest, error)
}

type gameRequestRepository struct {
	db *gorm.DB
}

func NewGameRequestRepository(db *gorm.DB) GameRequestRepository {
	return &gameRequestRepository{db: db}
}

func (r *gameRequestRepository) Create(gameRequest *models.GameRequest) (*models.GameRequest, error) {
	if err := r.db.Create(gameRequest).Error; err != nil {
		return nil, err
	}
	return gameRequest, nil
}

func (r *gameRequestRepository) GetByID(id uint) (*models.GameRequest, error) {
	var gameRequest models.GameRequest
	if err := r.db.Preload("Game").First(&gameRequest, id).Error; err != nil {
		return nil, err
	}
	return &gameRequest, nil
}

func (r *gameRequestRepository) Update(gameRequest *models.GameRequest) (*models.GameRequest, error) {
	if err := r.db.Save(gameRequest).Error; err != nil {
		return nil, err
	}
	return gameRequest, nil
}

func (r *gameRequestRepository) GetByUserAndGame(userID uint, gameID uint) (*models.GameRequest, error) {
	var gameRequest models.GameRequest
	if err := r.db.Where("user_id = ? AND game_id = ?", userID, gameID).First(&gameRequest).Error; err != nil {
		return nil, err
	}
	return &gameRequest, nil
}
