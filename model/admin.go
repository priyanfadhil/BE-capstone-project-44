package model

type Admin struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ID_Number int    `json:"id_number"`
	Password  string `json:"name"`
}
