package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	NIK      string `json:"nik"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
