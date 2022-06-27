package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID                   int                   `json:"id" gorm:"primaryKey"`
	IDNumber             int                   `json:"id_number"`
	Password             string                `json:"name"`
	VaccinationLocations []VaccinationLocation `gorm:"foreignKey:CreatedBY"`
}
