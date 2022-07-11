package model

import "gorm.io/gorm"

type FamilyMember struct {
	gorm.Model
	ID        int     `json:"id" gorm:"primaryKey"`
	UserID    int     `json:"user_id"`
	NIK       string  `json:"nik"`
	Hubungan  string  `json:"hubungan"`
	Name      string  `json:"name"`
	Birthdate string  `json:"birthdate" gorm:"type:Date"`
	Gender    string  `json:"gender"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	Booking   Booking `gorm:"foreignKey:FamilyID"`
}
