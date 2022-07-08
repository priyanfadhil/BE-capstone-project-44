package model

type VaccinationLocation struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	User_ID    string `json:"user_id" gorm:"foreignkey"`
	Created_By string `json:"created_by" gorm:"foreignkey"`
	Name       string `json:"name"`
	Alamat     string `json:"alamat"`
}
