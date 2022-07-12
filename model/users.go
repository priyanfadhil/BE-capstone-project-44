package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int            `json:"id" gorm:"primaryKey"`
	NIK          string         `json:"nik"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Password     string         `json:"password"`
	Address      string         `json:"address"`
	FamilyMember []FamilyMember `gorm:"foreignKey:UserID"`
}
