package model

import "gorm.io/gorm"

type Apotek struct {
	gorm.Model `json:"-"`
	ID         uint       `json:"id"`
	Logo       string     `json:"logo"`
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	Apotekers  []Apoteker `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"apoteker_lists"`
	Pegawais   []Pegawai  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pegawai_lists"`
}
