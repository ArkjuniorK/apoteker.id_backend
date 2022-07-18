package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Apotek struct {
	gorm.Model `gorm:"embeded"`
	UUID       uuid.UUID  `json:"uuid" gorm:"unique;index"`
	Logo       string     `json:"logo"`
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	Apotekers  []Apoteker `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"apoteker_lists"`
}

type ApotekCreateBody struct {
	UUID    uuid.UUID `json:"_id" gorm:"unique;index"`
	Logo    string    `json:"logo"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}
