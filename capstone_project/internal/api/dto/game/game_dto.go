package game

import "time"

type GameCreateDTO struct {
	Name           string    `json:"name" binding:"required" example:"Chess Tournament"`
	StartTime      time.Time `json:"startTime" binding:"required" example:"2023-04-01T10:00:00Z"`
	EndTime        time.Time `json:"endTime" binding:"required" example:"2023-04-01T18:00:00Z"`
	MaxMember      int       `json:"maxMember" binding:"required" example:"10"`
	GameCategoryID uint      `json:"gameCategoryId" binding:"required" example:"1"`
	RequiredSkills []uint    `json:"requiredSkills"`
}

type GameUpdateDTO struct {
	Name           string    `json:"name" example:"Updated Chess Tournament"`
	StartTime      time.Time `json:"startTime" example:"2023-04-01T11:00:00Z"`
	EndTime        time.Time `json:"endTime" example:"2023-04-01T19:00:00Z"`
	MaxMember      int       `json:"maxMember" example:"12"`
	RequiredSkills []uint    `json:"requiredSkills"`
	Categories     []uint    `json:"categories"`
}

type GameResponseDTO struct {
	ID             uint               `json:"id" example:"1"`
	Name           string             `json:"name" example:"Chess Tournament"`
	StartTime      time.Time          `json:"startTime" example:"2023-04-01T10:00:00Z"`
	EndTime        time.Time          `json:"endTime" example:"2023-04-01T18:00:00Z"`
	GameOwner      UserDTO            `json:"gameOwner"`
	MaxMember      int                `json:"maxMember" example:"10"`
	GameCategory   GameCategoryDTO    `json:"gameCategory"`
	RequiredSkills []RequiredSkillDTO `json:"requiredSkills"`
	GameRequests   []GameRequestDTO   `json:"gameRequests"`
	CreatedAt      time.Time          `json:"createdAt" example:"2023-03-30T12:00:00Z"`
	UpdatedAt      time.Time          `json:"updatedAt" example:"2023-03-30T12:00:00Z"`
}

type UserDTO struct {
	ID        uint   `json:"id" example:"1"`
	Username  string `json:"username" example:"john_doe"`
	AvatarURL string `json:"avatarUrl" example:"https://example.com/avatar.jpg"`
}

type GameCategoryDTO struct {
	ID       uint   `json:"id" example:"1"`
	Name     string `json:"name" example:"Strategy"`
	ImageUrl string `json:"imageUrl"`
}

type RequiredSkillDTO struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"Problem Solving"`
	Label string `json:"label" example:"Problem Solving"`
	Value string `json:"value" example:"Problem Solving"`
}

type GameRequestCreateDTO struct {
	GameID uint `json:"gameId" binding:"required" example:"1"`
}

type GameRequestUpdateDTO struct {
	Status string `json:"status" binding:"required" example:"accepted"`
}

type GameRequestResponseDTO struct {
	ID     uint   `json:"id" example:"1"`
	UserID uint   `json:"userId" example:"1"`
	GameID uint   `json:"gameId" example:"1"`
	Status string `json:"status" example:"pending"`
}

type GameRequestDTO struct {
	ID     uint    `json:"id" example:"1"`
	Status string  `json:"status" example:"pending"`
	User   UserDTO `json:"user"`
}

type RequiredSkillCreateDTO struct {
	Skills []string `json:"skills" binding:"required" example:['Problem Solving', 'Teamwork', 'Strategy']`
}
