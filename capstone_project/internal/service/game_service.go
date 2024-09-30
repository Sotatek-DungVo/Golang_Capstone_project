package service

import (
	dto "capstone_project/internal/api/dto/game"
	"capstone_project/internal/models"
	"capstone_project/internal/repository"
)

type GameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{repo: repo}
}

func (s *GameService) CreateGame(createDTO dto.GameCreateDTO, userID uint) (*dto.GameResponseDTO, error) {
	game := models.Game{
		Name:           createDTO.Name,
		StartTime:      createDTO.StartTime,
		EndTime:        createDTO.EndTime,
		MaxMember:      createDTO.MaxMember,
		GameOwnerID:    userID,
		GameCategoryID: createDTO.GameCategoryID,
	}

	if len(createDTO.RequiredSkills) > 0 {
		skills, err := s.repo.GetRequiredSkillsByIDs(createDTO.RequiredSkills)
		if err != nil {
			return nil, err
		}
		game.RequiredSkills = skills
	}

	createdGame, err := s.repo.CreateWithAssociations(&game)
	if err != nil {
		return nil, err
	}

	return mapGameToResponseDTO(createdGame), nil
}

func (s *GameService) GetGameByID(id uint) (*dto.GameResponseDTO, error) {
	game, err := s.repo.GetByIDWithDetails(id)
	if err != nil {
		return nil, err
	}

	return mapGameToResponseDTO(game), nil
}

func (s *GameService) UpdateGame(id uint, updateDTO dto.GameUpdateDTO) (*dto.GameResponseDTO, error) {
	game, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updateDTO.Name != "" {
		game.Name = updateDTO.Name
	}
	if !updateDTO.StartTime.IsZero() {
		game.StartTime = updateDTO.StartTime
	}
	if !updateDTO.EndTime.IsZero() {
		game.EndTime = updateDTO.EndTime
	}

	updatedGame, err := s.repo.Update(game)
	if err != nil {
		return nil, err
	}

	return mapGameToResponseDTO(updatedGame), nil
}

func (s *GameService) DeleteGame(id uint) error {
	return s.repo.Delete(id)
}

func (s *GameService) ListGames(page int, limit int) ([]*dto.GameResponseDTO, error) {
	games, err := s.repo.ListWithDetailsAndRequests(page, limit)
	if err != nil {
		return nil, err
	}

	var gameDTOs []*dto.GameResponseDTO
	for _, game := range games {
		gameDTOs = append(gameDTOs, mapGameToResponseDTO(game))
	}

	return gameDTOs, nil
}

func mapGameToResponseDTO(game *models.Game) *dto.GameResponseDTO {
	requiredSkills := make([]dto.RequiredSkillDTO, len(game.RequiredSkills))
	for i, skill := range game.RequiredSkills {
		requiredSkills[i] = dto.RequiredSkillDTO{
			ID:   skill.ID,
			Name: skill.Name,
		}
	}

	gameRequests := make([]dto.GameRequestDTO, len(game.GameRequests))
    for i, request := range game.GameRequests {
        gameRequests[i] = dto.GameRequestDTO{
            ID:     request.ID,
            Status: string(request.Status),
            User: dto.UserDTO{
                ID:        request.User.ID,
                Username:  request.User.Username,
                AvatarURL: request.User.AvatarUrl,
            },
        }
    }

	return &dto.GameResponseDTO{
		ID:        game.ID,
		Name:      game.Name,
		StartTime: game.StartTime,
		EndTime:   game.EndTime,
		GameOwner: dto.UserDTO{
			ID:        game.GameOwner.ID,
			Username:  game.GameOwner.Username,
			AvatarURL: game.GameOwner.AvatarUrl,
		},
		MaxMember: game.MaxMember,
		GameCategory: dto.GameCategoryDTO{
			ID:       game.GameCategory.ID,
			Name:     game.GameCategory.Name,
			ImageUrl: game.GameCategory.ImageUrl,
		},
		GameRequests: gameRequests,
		RequiredSkills: requiredSkills,
		CreatedAt:      game.CreatedAt,
		UpdatedAt:      game.UpdatedAt,
	}
}
