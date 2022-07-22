package model

import "gorm.io/gorm"

type Apoteker struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"user_name"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_picture"`
	Apoteks    Apotek `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"apotek"`
}

type ApotekerSerialize struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"user_name"`
	ProfilePic string `json:"profile_picture"`
	ApotekName string `json:"apotek_name"`
}
