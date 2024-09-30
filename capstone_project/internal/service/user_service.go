package service

import (
	dto "capstone_project/internal/api/dto/user"
	"capstone_project/internal/models"
	"capstone_project/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(page int, limit int, search map[string]string) ([]*dto.UserResponseDTO, error) {
	users, err := s.repo.List(page, limit, search)
	if err != nil {
		return nil, err
	}

	var userDTOs []*dto.UserResponseDTO
	for _, user := range users {
		userDTOs = append(userDTOs, mapUserToResponseDTO(user))
	}

	return userDTOs, nil
}

func mapUserToResponseDTO(user *models.User) *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Description: user.Description,
		AvatarUrl: user.AvatarUrl,
		Gender:      user.Gender,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

