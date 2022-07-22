package model

import "gorm.io/gorm"

type Apotek struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id"`
	Logo       string    `json:"logo"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	ApotekerID uint      `json:"apoteker_id"`
	Pegawais   []Pegawai `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pegawai_lists"`
}
