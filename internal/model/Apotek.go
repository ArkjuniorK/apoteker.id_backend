package model

import (
	"github.com/google/uuid"
)

type Apotek struct {
	// gorm.Model
	ID      uuid.UUID `gorm:"type:uuid"`
	Logo    string
	Name    string
	Address string
}

func GetApotek() {
}
