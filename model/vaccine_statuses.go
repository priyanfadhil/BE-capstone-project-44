package model

import "gorm.io/gorm"

type VaccineStatus struct {
	gorm.Model
	ID      int     `json:"id" gorm:"primaryKey"`
	Status  string  `json:"status"`
	Booking Booking `gorm:"foreignKey:StatusVaccineID"`
}
