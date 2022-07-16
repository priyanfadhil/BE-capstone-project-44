package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockVaccinationLocationSvc struct {
	Mock mock.Mock
}

func (m *MockVaccinationLocationSvc) CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error {
	arg := m.Mock.Called(vaccinationlocation)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockVaccinationLocationSvc) UpdateVaccinationLocation(id int, vaccinationlocation model.VaccinationLocation) error {
	return nil
}
func (m *MockVaccinationLocationSvc) GetAllVaccinationLocations() []model.VaccinationLocation {
	return []model.VaccinationLocation{}
}
func (m *MockVaccinationLocationSvc) GetVaccinationLocationByID(id int) (model.VaccinationLocation, error) {
	return model.VaccinationLocation{}, nil
}
func (m *MockVaccinationLocationSvc) DeleteVaccinationLocationByID(id int) error {
	return nil
}
