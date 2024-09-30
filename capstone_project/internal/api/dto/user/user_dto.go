package user

import (
	"capstone_project/internal/models"
	"time"
)

type UserResponseDTO struct {
	ID          uint          `json:"id"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Description string        `json:"description"`
	AvatarUrl   string        `json:"avatarUrl"`
	Gender      models.Gender `json:"gender"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
