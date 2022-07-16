package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type VaccineRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *VaccineRepositoryMysqlLayer) CreateVaccine(vaccine model.Vaccine) error {
	res := r.DB.Create(&vaccine)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *VaccineRepositoryMysqlLayer) GetAllVaccines() []model.Vaccine {
	vanncine := []model.Vaccine{}
	r.DB.Find(&vanncine)

	return vanncine
}

func (r *VaccineRepositoryMysqlLayer) GetOneVaccineByID(id int) (vaccine model.Vaccine, err error) {
	res := r.DB.Where("id = ?", id).Find(&vaccine)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *VaccineRepositoryMysqlLayer) UpdateOneVaccineByID(id int, vaccine model.Vaccine) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&vaccine)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *VaccineRepositoryMysqlLayer) DeleteVaccineByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Vaccine{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewVaccineMysqlRepository(db *gorm.DB) domain.AdapterVaccineRepository {
	return &VaccineRepositoryMysqlLayer{
		DB: db,
	}
}
