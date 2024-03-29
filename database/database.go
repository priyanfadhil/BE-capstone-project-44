package database

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(conectionString))
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.FamilyMember{})
	DB.AutoMigrate(&model.Booking{})
	DB.AutoMigrate(&model.Session{})
	DB.AutoMigrate(&model.VaccinationLocation{})
	DB.AutoMigrate(&model.Vaccine{})
	DB.AutoMigrate(&model.Admin{})

	return DB
}
