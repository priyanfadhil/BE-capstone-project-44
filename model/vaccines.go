package model

import "gorm.io/gorm"

type Vaccine struct {
	gorm.Model
	ID       int       `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Sessions []Session `gorm:"foreignKey:VaccineID"`
}
