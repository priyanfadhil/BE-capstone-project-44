package model

type VaccinationLocation struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	User_ID    int    `json:"user_id" gorm:"foreignkey"`
	Created_BY int    `json:"created_by" gorm:"foreignkey"`
	Name       string `json:"alamat"`
	Alamat     string `json:"kelurahan"`
}
