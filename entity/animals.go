package entity

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name      string `json:"name"`
	Type      string `json:"type"`
	Age       int    `json:"age"`
	IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
}
