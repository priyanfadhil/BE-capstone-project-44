package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockVaccineSvc struct {
	Mock mock.Mock
}

func (m *MockVaccineSvc) CreateVaccine(vaccine model.Vaccine) error {
	arg := m.Mock.Called(vaccine)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockVaccineSvc) UpdateVaccine(id int, vaccine model.Vaccine) error {
	return nil
}
func (m *MockVaccineSvc) GetAllVaccines() []model.Vaccine {
	return []model.Vaccine{}
}
func (m *MockVaccineSvc) GetVaccineByID(id int) (model.Vaccine, error) {
	return model.Vaccine{}, nil
}
func (m *MockVaccineSvc) DeleteVaccineByID(id int) error {
	return nil
}
