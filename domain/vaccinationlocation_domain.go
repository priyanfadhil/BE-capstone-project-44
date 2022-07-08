package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterVaccinationLocationRepository interface {
	CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error
	GetAllVaccinationLocation() []model.VaccinationLocation
	GetOneVaccinationLocationByID(id int) (vaccinationlocation model.VaccinationLocation, err error)
	GetOneVaccinationLocationByName(name string) (vaccinationlocation model.VaccinationLocation, err error)
	UpdateOneVaccinationLocationByID(id int, vaccinationlocation model.VaccinationLocation) error
	DeleteVaccinationLocationByID(id int) error
}

// Use Case
type AdapterVaccinationLocation interface {
	CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error
	UpdateVaccinationLocation(id, idToken int, vaccinationlocation model.VaccinationLocation) error
	GetAllVaccinationLocation() []model.VaccinationLocation
	GetVaccinationLocationByID(id int) (model.VaccinationLocation, error)
	DeleteVaccinationLocationByID(id int) error
}
