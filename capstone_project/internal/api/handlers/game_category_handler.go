package handlers

import (
	"capstone_project/internal/api/dto/category"
	"capstone_project/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GameCategoryHandler struct {
	service *service.GameCategoryService
}

func NewGameCategoryHandler(service *service.GameCategoryService) *GameCategoryHandler {
	return &GameCategoryHandler{service: service}
}

// ListGameCategories godoc
// @Summary List all game categories with pagination
// @Description Get a list of all game categories with pagination
// @Tags game-categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} category.GameCategoryResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /game-categories [get]
func (h *GameCategoryHandler) ListGameCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	categories, err := h.service.ListGameCategories(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// CreateGameCategory godoc
// @Summary Create a new game category
// @Description Create a new game category with the provided details
// @Tags game-categories
// @Accept json
// @Produce json
// @Param category body category.GameCategoryCreateDTO true "Game category details"
// @Success 201 {object} category.GameCategoryResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /game-categories [post]
func (h *GameCategoryHandler) CreateGameCategory(c *gin.Context) {
	var categoryDTO category.GameCategoryCreateDTO
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := h.service.CreateGameCategory(&categoryDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}
