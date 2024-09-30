package models

import (
	"gorm.io/gorm"
)

type GameCategory struct {
	gorm.Model
	Name     string `gorm:"unique;not null" json:"name"`
	ImageUrl string `json:"imageUrl"`
	Games    []Game `gorm:"foreignKey:GameCategoryID" json:"games,omitempty"`
}
