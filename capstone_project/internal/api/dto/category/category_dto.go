package category

import "time"

type GameCategoryResponseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GameCategoryCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
}
