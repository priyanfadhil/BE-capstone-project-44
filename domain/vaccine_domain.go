package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterVaccineRepository interface {
	CreateVaccine(vaccine model.Vaccine) error
	GetAllVaccine() []model.Vaccine
	GetOneVaccineByID(id int) (vaccine model.Vaccine, err error)
	GetOneVaccineByName(name string) (vaccine model.Vaccine, err error)
	UpdateOneVaccineByID(id int, vaccine model.Vaccine) error
	DeleteVaccineByID(id int) error
}

// Use Case
type AdapterVaccine interface {
	CreateVaccine(vaccine model.Vaccine) error
	UpdateVaccine(id, idToken int, vaccine model.Vaccine) error
	GetAllVaccine() []model.Vaccine
	GetVaccineByID(id int) (model.Vaccine, error)
	DeleteVaccineByID(id int) error
}
