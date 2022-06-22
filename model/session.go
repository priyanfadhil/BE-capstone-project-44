package model
import (
	"time"
)

type Session struct {	
	ID         		int 		`json:"id" gorm:"primaryKey"`
	Location_ID   	int 		`json:"location_id" gorm:"foreignkey"`
	Vaccine_ID 		int 		`json:"vaccine_id" gorm:"foreignkey"`
	Tanggal    		time.Time 	`json:"tanggal"`	
	Status     		bool 		`json:"status"`
	Stock_Vaccine   int  		`json:"stock_vaccine"`			
}