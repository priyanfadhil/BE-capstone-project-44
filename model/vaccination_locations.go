package model

import "gorm.io/gorm"

type VaccinationLocation struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	CreatedBY int       `json:"created_by"`
	Name      string    `json:"name"`
	Alamat    string    `json:"alamat"`
	Sessions  []Session `gorm:"foreignKey:LocationID"`
}
