package model

type Apotek struct {
	// gorm.Model
	ID      uint
	Logo    string
	Name    string
	Address string
}
