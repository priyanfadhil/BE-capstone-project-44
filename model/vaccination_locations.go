package model

import "gorm.io/gorm"

type VaccinationLocation struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedBY int       `json:"created_by"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Sessions  []Session `gorm:"foreignKey:LocationID"`
}
