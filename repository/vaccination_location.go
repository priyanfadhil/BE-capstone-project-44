package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type vaccinationlocationRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *vaccinationlocationRepositoryMysqlLayer) CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error {
	res := r.DB.Create(&vaccinationlocation)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *vaccinationlocationRepositoryMysqlLayer) GetAllVaccinationLocations() []model.VaccinationLocation {
	vaccinationlocations := []model.VaccinationLocation{}
	r.DB.Find(&vaccinationlocations)

	return vaccinationlocations
}

func (r *vaccinationlocationRepositoryMysqlLayer) GetOneVaccinationLocationByID(id int) (vaccinationlocation model.VaccinationLocation, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccinationlocation)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *vaccinationlocationRepositoryMysqlLayer) UpdateOneVaccinationLocationByID(id int, vaccinationlocation model.VaccinationLocation) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccinationlocation)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *vaccinationlocationRepositoryMysqlLayer) DeleteVaccinationLocationByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.VaccinationLocation{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func VaccinationLocationMysqlRepository(db *gorm.DB) domain.AdapterVaccinationLocationRepository {
	return &vaccinationlocationRepositoryMysqlLayer{
		DB: db,
	}
}
