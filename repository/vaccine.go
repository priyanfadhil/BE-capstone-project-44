package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type vaccineRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *vaccineRepositoryMysqlLayer) CreateVaccine(vaccine model.Vaccine) error {
	res := r.DB.Create(&vaccine)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *vaccineRepositoryMysqlLayer) GetAllVaccine() []model.Vaccine {
	vaccine := []model.Vaccine{}
	r.DB.Find(&vaccine)
	//r.DB.Model(&model.Vaccine{}).Association("rents").Find(&vaccine)

	return vaccine
}

func (r *vaccineRepositoryMysqlLayer) GetOneVaccineByID(id int) (vaccine model.Vaccine, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccine)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *vaccineRepositoryMysqlLayer) GetOneVaccineByName(name string) (vaccine model.Vaccine, err error) {
	res := r.DB.Where("name = ?", name).Find(&vaccine)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return vaccine, err
}

func (r *vaccineRepositoryMysqlLayer) UpdateOneVaccineByID(id int, vaccine model.Vaccine) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccine)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *vaccineRepositoryMysqlLayer) DeleteVaccineByID(id int) error {
	res := r.DB.Delete(&model.Vaccine{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func VaccineMysqlRepository(db *gorm.DB) domain.AdapterVaccineRepository {
	return &vaccineRepositoryMysqlLayer{
		DB: db,
	}
}
