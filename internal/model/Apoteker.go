package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Apoteker struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"_id" gorm:"unique;index"`
	FullName   string    `json:"full_name"`
	Username   string    `json:"user_name"`
	Password   string    `json:"password"`
	ProfilePic string    `json:"profile_picture"`
	ApotekID   uint      `json:"apotek_id"`
}

type ApotekerSerialize struct {
	ID         uuid.UUID `json:"_id"`
	FullName   string    `json:"full_name"`
	Username   string    `json:"user_name"`
	ProfilePic string    `json:"profile_picture"`
	ApotekName string    `json:"apotek_name"`
}
