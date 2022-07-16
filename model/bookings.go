package model

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID            int    `json:"id" gorm:"primaryKey"`
	FamilyID      int    `json:"family_id" gorm:"foreignkey"`
	SessionID     int    `json:"session_id" gorm:"foreignkey"`
	VaccineStatus int    `json:"vaccine_status"`
	TicketNumber  string `json:"ticket_number"`
	Done          bool   `json:"done"`
}
