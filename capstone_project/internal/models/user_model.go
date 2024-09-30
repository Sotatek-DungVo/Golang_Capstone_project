package models

import (
	"gorm.io/gorm"
)

type Gender string

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique;not null" json:"username"`
	Email       string `gorm:"unique;not null" json:"email"`
	AvatarUrl   string `json:"avatarUrl"`
	Password    string `gorm:"not null" json:"-"`
	Description string `json:"description"`
	IsEnabled   bool   `gorm:"default:true" json:"isEnabled"`
	Gender      Gender `gorm:"type:varchar(20);not null" json:"gender"`
	OwnedGames   []Game        `gorm:"foreignKey:GameOwnerID" json:"ownedGames,omitempty"`
	GameRequests []GameRequest `gorm:"foreignKey:UserID" json:"gameRequests,omitempty"`
}
