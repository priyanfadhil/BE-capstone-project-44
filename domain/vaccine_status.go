package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterVaccineStatusRepository interface {
	CreateVaccineStatus(VaccineStatus model.VaccineStatus) error
	GetAllVaccineStatuses() []model.VaccineStatus
	GetOneVaccineStatusByID(id int) (VaccineStatus model.VaccineStatus, err error)
	UpdateOneVaccineStatusByID(id int, VaccineStatus model.VaccineStatus) error
	DeleteVaccineStatusByID(id int) error
}

// Use Case
type AdapterVaccineStatus interface {
	CreateVaccineStatus(VaccineStatus model.VaccineStatus) error
	UpdateVaccineStatus(id int, VaccineStatus model.VaccineStatus) error
	GetAllVaccineStatuses() []model.VaccineStatus
	GetVaccineStatusByID(id int) (model.VaccineStatus, error)
	DeleteVaccineStatusByID(id int) error
}
