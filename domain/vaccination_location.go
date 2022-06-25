package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterVaccinationLocationRepository interface {
	CreateVaccinationLocation(VaccinationLocation model.VaccinationLocation) error
	GetAllVaccinationLocations() []model.VaccinationLocation
	GetOneVaccinationLocationByID(id int) (VaccinationLocation model.VaccinationLocation, err error)
	UpdateOneVaccinationLocationByID(id int, VaccinationLocation model.VaccinationLocation) error
	DeleteVaccinationLocationByID(id int) error
}

// Use Case
type AdapterVaccinationLocation interface {
	CreateVaccinationLocation(VaccinationLocation model.VaccinationLocation) error
	UpdateVaccinationLocation(id int, VaccinationLocation model.VaccinationLocation) error
	GetAllVaccinationLocations() []model.VaccinationLocation
	GetVaccinationLocationByID(id int) (model.VaccinationLocation, error)
	DeleteVaccinationLocationByID(id int) error
}
