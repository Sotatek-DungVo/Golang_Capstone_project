package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name      string    `gorm:"unique;not null" json:"name"`
	MaxMember int       `json:"maxMember"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`

	GameOwnerID uint `gorm:"not null" json:"gameOwnerId"`
	GameOwner   User `gorm:"foreignKey:GameOwnerID" json:"gameOwner"`

	GameCategoryID uint         `gorm:"not null" json:"gameCategoryId"`
	GameCategory   GameCategory `gorm:"foreignKey:GameCategoryID" json:"gameCategory"`

	RequiredSkills []RequiredSkill `gorm:"many2many:game_required_skills;" json:"requiredSkills,omitempty"`
	GameRequests   []GameRequest   `gorm:"foreignKey:GameID" json:"gameRequests,omitempty"`
}
