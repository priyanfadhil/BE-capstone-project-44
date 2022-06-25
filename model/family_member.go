package model

type FamilyMember struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	NIK    string `json:"nik"`
	Name   string `json:"name"`
	UserID int    `json:"userid"`
}
