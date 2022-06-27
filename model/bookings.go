package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	ID              int  `json:"id" gorm:"primaryKey"`
	FamilyID        int  `json:"family_id" gorm:"foreignkey"`
	SessionID       int  `json:"session_id" gorm:"foreignkey"`
	StatusVaccineID int  `json:"status_vaccine_id" gorm:"foreignkey"`
	Status          bool `json:"status"`
}
