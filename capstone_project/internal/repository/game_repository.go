package repository

import (
	"capstone_project/internal/models"
	"gorm.io/gorm"
)

type GameRepository interface {
	GetByID(id uint) (*models.Game, error)
	Update(game *models.Game) (*models.Game, error)
	Delete(id uint) error
	List(page int, limit int) ([]*models.Game, error)
	ListWithDetails(page int, limit int) ([]*models.Game, error)
	ListWithDetailsAndRequests(page int, limit int) ([]*models.Game, error)
	GetByIDWithDetails(id uint) (*models.Game, error)
	CreateWithAssociations(game *models.Game) (*models.Game, error)
	GetRequiredSkillsByIDs(ids []uint) ([]models.RequiredSkill, error)
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) GetByID(id uint) (*models.Game, error) {
	var game models.Game
	if err := r.db.First(&game, id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}

func (r *gameRepository) Update(game *models.Game) (*models.Game, error) {
	if err := r.db.Save(game).Error; err != nil {
		return nil, err
	}
	return game, nil
}

func (r *gameRepository) Delete(id uint) error {
	return r.db.Delete(&models.Game{}, id).Error
}

func (r *gameRepository) List(page int, limit int) ([]*models.Game, error) {
	var games []*models.Game
	offset := (page - 1) * limit
	if err := r.db.Offset(offset).Limit(limit).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (r *gameRepository) ListWithDetails(page int, limit int) ([]*models.Game, error) {
	var games []*models.Game
	offset := (page - 1) * limit
	if err := r.db.Preload("GameOwner").Preload("GameCategory").Preload("RequiredSkills").Offset(offset).Limit(limit).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (r *gameRepository) ListWithDetailsAndRequests(page int, limit int) ([]*models.Game, error) {
	var games []*models.Game
	offset := (page - 1) * limit
	if err := r.db.Preload("GameOwner").
		Preload("GameCategory").
		Preload("RequiredSkills").
		Preload("GameRequests").
		Preload("GameRequests.User").
		Offset(offset).Limit(limit).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (r *gameRepository) GetByIDWithDetails(id uint) (*models.Game, error) {
	var game models.Game
	if err := r.db.Preload("GameOwner").
		Preload("GameCategory").
		Preload("RequiredSkills").
		Preload("GameRequests").
		Preload("GameRequests.User").
		First(&game, id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}

func (r *gameRepository) CreateWithAssociations(game *models.Game) (*models.Game, error) {
	if err := r.db.Create(game).Error; err != nil {
		return nil, err
	}

	var createdGame models.Game
	if err := r.db.Preload("GameOwner").Preload("GameCategory").Preload("RequiredSkills").First(&createdGame, game.ID).Error; err != nil {
		return nil, err
	}

	return &createdGame, nil
}

func (r *gameRepository) GetRequiredSkillsByIDs(ids []uint) ([]models.RequiredSkill, error) {
	var skills []models.RequiredSkill
	if err := r.db.Where("id IN ?", ids).Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}
