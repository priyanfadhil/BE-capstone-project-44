package model

type UserAddress struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	User_ID   int    `json:"user_id" gorm:"foreignkey"`
	Alamat    string `json:"alamat"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kota      string `json:"kota"`
	Provinsi  string `json:"provinsi"`
}
