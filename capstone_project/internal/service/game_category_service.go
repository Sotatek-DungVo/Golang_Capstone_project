package service

import (
	dto "capstone_project/internal/api/dto/category"
	"capstone_project/internal/models"
	"capstone_project/internal/repository"
)

type GameCategoryService struct {
	repo *repository.GameCategoryRepository
}

func NewGameCategoryService(repo *repository.GameCategoryRepository) *GameCategoryService {
	return &GameCategoryService{repo: repo}
}

func (s *GameCategoryService) ListGameCategories(page int, limit int) ([]*dto.GameCategoryResponseDTO, error) {
	categories, err := s.repo.List(page, limit)
	if err != nil {
		return nil, err
	}

	var categoryDTOs []*dto.GameCategoryResponseDTO
	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, mapGameCategoryToResponseDTO(category))
	}

	return categoryDTOs, nil
}

func (s *GameCategoryService) CreateGameCategory(categoryDTO *dto.GameCategoryCreateDTO) (*dto.GameCategoryResponseDTO, error) {
	category := &models.GameCategory{
		Name:     categoryDTO.Name,
		ImageUrl: categoryDTO.ImageUrl,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}

	return mapGameCategoryToResponseDTO(category), nil
}

func mapGameCategoryToResponseDTO(category *models.GameCategory) *dto.GameCategoryResponseDTO {
	return &dto.GameCategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		ImageUrl:  category.ImageUrl,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
