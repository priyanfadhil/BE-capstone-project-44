package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type vaccineStatusRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *vaccineStatusRepositoryMysqlLayer) CreateVaccineStatus(vaccinestatus model.VaccineStatus) error {
	res := r.DB.Create(&vaccinestatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *vaccineStatusRepositoryMysqlLayer) GetAllVaccineStatus() []model.VaccineStatus {
	vaccinestatus := []model.VaccineStatus{}
	r.DB.Find(&vaccinestatus)
	//r.DB.Model(&model.VaccineStatus{}).Association("rents").Find(&vaccinestatus)

	return vaccinestatus
}

func (r *vaccineStatusRepositoryMysqlLayer) GetOneVaccineStatusByID(id int) (vaccinestatus model.VaccineStatus, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccinestatus)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *vaccineStatusRepositoryMysqlLayer) GetOneVaccineStatusByName(name string) (vaccinestatus model.VaccineStatus, err error) {
	res := r.DB.Where("name = ?", name).Find(&vaccinestatus)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return vaccinestatus, err
}

func (r *vaccineStatusRepositoryMysqlLayer) UpdateOneVaccineStatusByID(id int, vaccinestatus model.VaccineStatus) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccinestatus)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *vaccineStatusRepositoryMysqlLayer) DeleteVaccineStatusByID(id int) error {
	res := r.DB.Delete(&model.VaccineStatus{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func VaccineStatusMysqlRepository(db *gorm.DB) domain.AdapterVaccineStatusRepository {
	return &vaccineStatusRepositoryMysqlLayer{
		DB: db,
	}
}
