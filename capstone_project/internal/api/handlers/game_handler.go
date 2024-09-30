package handlers

import (
	"capstone_project/internal/api/dto/game"
	"capstone_project/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type GameHandler struct {
	gameService *service.GameService
}

func NewGameHandler(gameService *service.GameService) *GameHandler {
	return &GameHandler{gameService: gameService}
}

// CreateGame godoc
// @Summary Create a new game
// @Description Create a new game with the provided details
// @Tags games
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param game body game.GameCreateDTO true "Game details"
// @Success 201 {object} game.GameResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /games [post]
func (h *GameHandler) CreateGame(c *gin.Context) {
	var createDTO game.GameCreateDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	createdGame, err := h.gameService.CreateGame(createDTO, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdGame)
}

// GetGame godoc
// @Summary Get a game by ID
// @Description Get details of a game by its ID, including game category, required skills, game owner, and game requests
// @Tags games
// @Accept json
// @Produce json
// @Param id path int true "Game ID"
// @Success 200 {object} game.GameResponseDTO
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /games/{id} [get]
func (h *GameHandler) GetGame(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid game ID"})
		return
	}

	game, err := h.gameService.GetGameByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// UpdateGame godoc
// @Summary Update a game
// @Description Update a game with the provided details
// @Tags games
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param id path int true "Game ID"
// @Param game body game.GameUpdateDTO true "Game details to update"
// @Success 200 {object} game.GameResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /games/{id} [put]
func (h *GameHandler) UpdateGame(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid game ID"})
		return
	}

	var updateDTO game.GameUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	// Check if the user owns the game
	game, err := h.gameService.GetGameByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Game not found"})
		return
	}
	if game.GameOwner.ID != userId.(uint) {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "You don't have permission to update this game"})
		return
	}

	updatedGame, err := h.gameService.UpdateGame(uint(id), updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedGame)
}

// DeleteGame godoc
// @Summary Delete a game
// @Description Delete a game by its ID
// @Tags games
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param id path int true "Game ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /games/{id} [delete]
func (h *GameHandler) DeleteGame(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid game ID"})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	// Add a check to ensure the user owns the game
	game, err := h.gameService.GetGameByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Game not found"})
		return
	}
	if game.GameOwner.ID != userId.(uint) {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "You don't have permission to delete this game"})
		return
	}

	if err := h.gameService.DeleteGame(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListGames godoc
// @Summary List all games with pagination
// @Description Get a list of all games with pagination, including game category details, required skills, and game requests
// @Tags games
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} game.GameResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /games [get]
func (h *GameHandler) ListGames(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid limit number"})
		return
	}

	games, err := h.gameService.ListGames(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
