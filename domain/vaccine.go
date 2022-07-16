package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterVaccineRepository interface {
	CreateVaccine(Vaccine model.Vaccine) error
	GetAllVaccines() []model.Vaccine
	GetOneVaccineByID(id int) (Vaccine model.Vaccine, err error)
	UpdateOneVaccineByID(id int, Vaccine model.Vaccine) error
	DeleteVaccineByID(id int) error
}

// Use Case
type AdapterVaccine interface {
	CreateVaccine(Vaccine model.Vaccine) error
	UpdateVaccine(id int, Vaccine model.Vaccine) error
	GetAllVaccines() []model.Vaccine
	GetVaccineByID(id int) (model.Vaccine, error)
	DeleteVaccineByID(id int) error
}
