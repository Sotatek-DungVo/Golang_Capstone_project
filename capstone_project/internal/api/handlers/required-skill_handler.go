package handlers

import (
	"capstone_project/internal/api/dto/game"
	"capstone_project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequiredSkillHandler struct {
	requiredSkillService *service.RequiredSkillService
}

func NewRequiredSkillHandler(requiredSkillService *service.RequiredSkillService) *RequiredSkillHandler {
	return &RequiredSkillHandler{requiredSkillService: requiredSkillService}
}

// CreateRequiredSkill godoc
// @Summary Create a new required skill
// @Description Create a new required skill (requires user login)
// @Tags required-skills
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param skill body game.RequiredSkillCreateDTO true "Required skill details"
// @Success 201 {object} game.RequiredSkillDTO
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /required-skills [post]
func (h *RequiredSkillHandler) CreateRequiredSkill(c *gin.Context) {
	var createDTO game.RequiredSkillCreateDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	skill, err := h.requiredSkillService.CreateRequiredSkill(createDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, skill)
}

// ListRequiredSkills godoc
// @Summary List all required skills
// @Description Get a list of all required skills from the database
// @Tags required-skills
// @Accept json
// @Produce json
// @Success 200 {array} game.RequiredSkillDTO
// @Failure 500 {object} ErrorResponse
// @Router /required-skills [get]
func (h *RequiredSkillHandler) ListRequiredSkills(c *gin.Context) {
	skills, err := h.requiredSkillService.ListRequiredSkills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, skills)
}

// CreateRequiredSkills godoc
// @Summary Create multiple new required skills
// @Description Create multiple new required skills (requires user login)
// @Tags required-skills
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param skills body game.RequiredSkillCreateDTO true "Required skills details"
// @Success 201 {array} game.RequiredSkillDTO
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /required-skills [post]
func (h *RequiredSkillHandler) CreateRequiredSkills(c *gin.Context) {
	var createDTO game.RequiredSkillCreateDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not authenticated"})
		return
	}

	skills, err := h.requiredSkillService.CreateRequiredSkills(createDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, skills)
}
