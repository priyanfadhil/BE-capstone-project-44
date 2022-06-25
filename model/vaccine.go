package model

type Vaccine struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
