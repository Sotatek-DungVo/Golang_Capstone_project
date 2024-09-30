package models

import (
	"gorm.io/gorm"
)

type GameRequest struct {
	gorm.Model
	UserID uint              `gorm:"not null" json:"userId"`
	GameID uint              `gorm:"not null" json:"gameId"`
	Status GameRequestStatus `gorm:"type:varchar(20);not null" json:"status"`

	User User `gorm:"foreignKey:UserID" json:"user"`
	Game Game `gorm:"foreignKey:GameID" json:"game"`
}
