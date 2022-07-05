package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID                   int                   `json:"id" gorm:"primaryKey"`
	Email                string                `json:"email"`
	Password             string                `json:"password"`
	VaccinationLocations []VaccinationLocation `gorm:"foreignKey:CreatedBY"`
}
