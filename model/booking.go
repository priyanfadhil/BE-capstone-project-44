package model

type Booking struct {
	ID		 			int 	`json:"id" gorm:"primaryKey"`
	Family_ID       	int     `json:"family_id" gorm:"foreignkey"`
	Session_ID      	int 	`json:"session_id" gorm:"foreignkey"`
	Status_Vaccine_ID   int 	`json:"status_vaccine_id" gorm:"foreignkey"`
	Status    			bool 	`json:"status"`
}