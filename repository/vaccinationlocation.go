package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type VaccinationLocationRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *VaccinationLocationRepositoryMysqlLayer) CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error {
	res := r.DB.Create(&vaccinationlocation)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *VaccinationLocationRepositoryMysqlLayer) GetAllVaccinationLocation() []model.VaccinationLocation {
	vaccinationlocation := []model.VaccinationLocation{}
	r.DB.Find(&vaccinationlocation)
	//r.DB.Model(&model.VaccinationLocation{}).Association("rents").Find(&vaccinationlocation)

	return vaccinationlocation
}

func (r *VaccinationLocationRepositoryMysqlLayer) GetOneVaccinationLocationByID(id int) (vaccinationlocation model.VaccinationLocation, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccinationlocation)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return vaccinationlocation
}

func (r *VaccinationLocationRepositoryMysqlLayer) GetOneVaccinationLocationByName(name string) (vaccinationlocation model.VaccinationLocation, err error) {
	res := r.DB.Where("name = ?", name).Find(&vaccinationlocation)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return vaccinationlocation, err
}

func (r *VaccinationLocationRepositoryMysqlLayer) UpdateOneVaccinationLocationByID(id int, vaccinationlocation model.VaccinationLocation) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccinationlocation)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *VaccinationLocationRepositoryMysqlLayer) DeleteVaccinationLocationByID(id int) error {
	res := r.DB.Delete(&model.VaccinationLocation{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func VaccinationLocationMysqlRepository(db *gorm.DB) domain.AdapterVaccinationLocationRepository {
	return &VaccinationLocationRepositoryMysqlLayer{
		DB: db,
	}
}
