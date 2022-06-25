package model

type VaccineStatus struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
