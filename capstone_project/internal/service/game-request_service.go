package service

import (
	"capstone_project/internal/models"
	"capstone_project/internal/repository"
	"capstone_project/internal/api/dto/game"
)

type GameRequestService struct {
	repo repository.GameRequestRepository
}

func NewGameRequestService(repo repository.GameRequestRepository) *GameRequestService {
	return &GameRequestService{repo: repo}
}

func (s *GameRequestService) CreateGameRequest(createDTO game.GameRequestCreateDTO, userID uint) (*game.GameRequestResponseDTO, error) {
	gameRequest := models.GameRequest{
		UserID: userID,
		GameID: createDTO.GameID,
		Status: models.Pending,
	}

	createdGameRequest, err := s.repo.Create(&gameRequest)
	if err != nil {
		return nil, err
	}

	return mapGameRequestToResponseDTO(createdGameRequest), nil
}

func (s *GameRequestService) GetGameRequestByID(id uint) (*game.GameRequestResponseDTO, error) {
	gameRequest, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mapGameRequestToResponseDTO(gameRequest), nil
}

func (s *GameRequestService) UpdateGameRequest(id uint, updateDTO game.GameRequestUpdateDTO) (*game.GameRequestResponseDTO, error) {
	gameRequest, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updateDTO.Status != "" {
		gameRequest.Status = models.GameRequestStatus(updateDTO.Status)
	}

	updatedGameRequest, err := s.repo.Update(gameRequest)
	if err != nil {
		return nil, err
	}

	return mapGameRequestToResponseDTO(updatedGameRequest), nil
}

func mapGameRequestToResponseDTO(gameRequest *models.GameRequest) *game.GameRequestResponseDTO {
	return &game.GameRequestResponseDTO{
		ID:     gameRequest.ID,
		UserID: gameRequest.UserID,
		GameID: gameRequest.GameID,
		Status: string(gameRequest.Status),
	}
}


func (s *GameRequestService) GetGameRequestByUserAndGame(userID uint, gameID uint) (*game.GameRequestResponseDTO, error) {
	gameRequest, err := s.repo.GetByUserAndGame(userID, gameID)
	if err != nil {
		return nil, err
	}
	return mapGameRequestToResponseDTO(gameRequest), nil
}