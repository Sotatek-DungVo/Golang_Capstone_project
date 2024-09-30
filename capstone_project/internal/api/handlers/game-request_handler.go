package handlers

import (
	"capstone_project/internal/api/dto/game"
	"capstone_project/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type GameRequestHandler struct {
	gameRequestService *service.GameRequestService
	gameService        *service.GameService
}

func NewGameRequestHandler(gameRequestService *service.GameRequestService, gameService *service.GameService) *GameRequestHandler {
	return &GameRequestHandler{gameRequestService: gameRequestService, gameService: gameService}
}

// CreateGameRequest godoc
// @Summary Create a new game request
// @Description Create a new game request with status pending
// @Tags game-requests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param gameRequest body game.GameRequestCreateDTO true "Game request details"
// @Success 201 {object} game.GameRequestResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /game-requests [post]
func (h *GameRequestHandler) CreateGameRequest(c *gin.Context) {
	var createDTO game.GameRequestCreateDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	gameDetails, err := h.gameService.GetGameByID(createDTO.GameID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Game not found"})
		return
	}

	existingRequest, err := h.gameRequestService.GetGameRequestByUserAndGame(userId.(uint), createDTO.GameID)
	if err == nil && existingRequest != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "You have already requested to join this game"})
		return
	}

	if gameDetails.GameOwner.ID == userId.(uint) {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "You already join your game"})
		return
	}

	if time.Now().After(gameDetails.StartTime) {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "Game is expired"})
		return
	}

	if len(gameDetails.GameRequests) >= gameDetails.MaxMember-1 {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "The game is full"})
		return
	}

	gameRequest, err := h.gameRequestService.CreateGameRequest(createDTO, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gameRequest)
}

// UpdateGameRequest godoc
// @Summary Update a game request
// @Description Approve or reject a game request
// @Tags game-requests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param id path int true "Game Request ID"
// @Param gameRequest body game.GameRequestUpdateDTO true "Game request details to update"
// @Success 200 {object} game.GameRequestResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /game-requests/{id} [put]
func (h *GameRequestHandler) UpdateGameRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid game request ID"})
		return
	}

	var updateDTO game.GameRequestUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	gameRequest, err := h.gameRequestService.GetGameRequestByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Game request not found"})
		return
	}

	fmt.Println(userId)
	fmt.Println(gameRequest.UserID)
	if gameRequest.UserID != userId.(uint) {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "You don't have permission to update this game request"})
		return
	}

	updatedGameRequest, err := h.gameRequestService.UpdateGameRequest(uint(id), updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedGameRequest)
}
