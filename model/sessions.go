package model

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID           int     `json:"id" gorm:"primaryKey"`
	LocationID   int     `json:"location_id" gorm:"foreignkey"`
	VaccineID    int     `json:"vaccine_id" gorm:"foreignkey"`
	Name         string  `json:"name"`
	Date         string  `json:"date" gorm:"type:Date"`
	Start        string  `json:"start" gorm:"type:Datetime" time_format:"2022-06-01 10:00"`
	End          string  `json:"end" gorm:"type:Datetime" time_format:"2022-06-01 10:00"`
	Status       *bool   `json:"status"`
	StockVaccine *int64  `json:"stock_vaccine"`
	Booking      Booking `gorm:"foreignKey:SessionID"`
}
