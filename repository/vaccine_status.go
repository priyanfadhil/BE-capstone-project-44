package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type VaccineStatusRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *VaccineStatusRepositoryMysqlLayer) CreateVaccineStatus(vaccinestatus model.VaccineStatus) error {
	res := r.DB.Create(&vaccinestatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *VaccineStatusRepositoryMysqlLayer) GetAllVaccineStatus() []model.VaccineStatus {
	vanncinestatus := []model.VaccineStatus{}
	r.DB.Find(&vanncinestatus)
	//r.DB.Model(&model.User{}).Association("rents").Find(&users)

	return vanncinestatus
}

func (r *VaccineStatusRepositoryMysqlLayer) GetOneVaccineStatusByID(id int) (vaccinestatus model.VaccineStatus, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccinestatus)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *VaccineStatusRepositoryMysqlLayer) UpdateOneVaccineStatusByID(id int, vaccinestatus model.VaccineStatus) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccinestatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *VaccineStatusRepositoryMysqlLayer) DeleteVaccineStatusByID(id int) error {
	res := r.DB.Delete(&model.VaccineStatus{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewVaccineStatusMysqlRepository(db *gorm.DB) domain.AdapterVaccineStatusRepository {
	return &VaccineStatusRepositoryMysqlLayer{
		DB: db,
	}
}
