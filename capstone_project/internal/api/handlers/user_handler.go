package handlers

import (
	"capstone_project/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// ListUsers godoc
// @Summary List all users with pagination and search
// @Description Get a list of all users with pagination and search by username and gender
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param username query string false "Search by username"
// @Param gender query string false "Search by gender (MALE or FEMALE)"
// @Success 200 {array} user.UserResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /players [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
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

	search := make(map[string]string)
	if name := c.Query("name"); name != "" {
		search["name"] = name
	}
	if username := c.Query("username"); username != "" {
		search["username"] = username
	}
	if gender := c.Query("gender"); gender != "" {
		search["gender"] = gender
	}

	users, err := h.userService.ListUsers(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
