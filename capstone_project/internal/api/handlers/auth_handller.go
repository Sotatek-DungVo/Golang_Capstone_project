package handlers

import (
	_ "capstone_project/docs"
	dto "capstone_project/internal/api/dto/auth"
	"capstone_project/internal/models"
	"capstone_project/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterDTO true "User registration information"
// @Success 201 {object} dto.RegisterResponseDTO
// @Router /auth/register [post]
func (h *AuthHandler) CreateUser(c *gin.Context) {
	var registerDTO dto.RegisterDTO

	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Username:    registerDTO.Username,
		Email:       registerDTO.Email,
		Password:    registerDTO.Password,
		Description: registerDTO.Description,
		Gender:      models.Gender(registerDTO.Gender),
		AvatarUrl:   registerDTO.AvatarUrl,
	}

	if err := h.service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.RegisterResponseDTO{
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusCreated, response)
}

// LoginUser godoc
// @Summary Login user
// @Description Login with username/email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginDTO true "Login credentials"
// @Success 200 {object} dto.LoginResponseDTO
// @Router /auth/login [post]
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var loginDTO dto.LoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.service.LoginUser(loginDTO.Identifier, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	response := dto.LoginResponseDTO{
		Token:     token,
		Username:  user.Username,
		AvatarUrl: user.AvatarUrl,
	}

	c.JSON(http.StatusOK, response)
}
