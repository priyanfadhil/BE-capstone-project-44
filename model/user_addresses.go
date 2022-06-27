package model

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	ID                   int                   `json:"id" gorm:"primaryKey"`
	UserID               int                   `json:"user_id"`
	Alamat               string                `json:"alamat"`
	Kelurahan            string                `json:"kelurahan"`
	Kecamatan            string                `json:"kecamatan"`
	Kota                 string                `json:"kota"`
	Provinsi             string                `json:"provinsi"`
	VaccinationLocations []VaccinationLocation `gorm:"foreignKey:UserID"`
}
