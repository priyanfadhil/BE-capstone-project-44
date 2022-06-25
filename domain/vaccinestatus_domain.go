package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterVaccineStatusRepository interface {
	CreateVaccineStatus(vaccinestatus model.VaccineStatus) error
	GetAllVaccineStatus() []model.VaccineStatus
	GetOneVaccineStatusByID(id int) (vaccinestatus model.VaccineStatus, err error)
	GetOneVaccineStatusByName(name string) (vaccinestatus model.VaccineStatus, err error)
	UpdateOneVaccineStatusByID(id int, vaccinestatus model.VaccineStatus) error
	DeleteVaccineStatusByID(id int) error
}

// Use Case
type AdapterVaccineStatus interface {
	CreateVaccineStatus(vaccinestatus model.VaccineStatus) error
	UpdateVaccineStatus(id, idToken int, vaccinestatus model.VaccineStatus) error
	GetAllVaccineStatus() []model.VaccineStatus
	GetVaccineStatusByID(id int) (model.VaccineStatus, error)
	DeleteVaccineStatusByID(id int) error
}
