package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockVaccineStatusSvc struct {
	Mock mock.Mock
}

func (m *MockVaccineStatusSvc) CreateVaccineStatus(vaccinestatus model.VaccineStatus) error {
	arg := m.Mock.Called(vaccinestatus)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockVaccineStatusSvc) UpdateVaccineStatus(id int, vaccinestatus model.VaccineStatus) error {
	return nil
}
func (m *MockVaccineStatusSvc) GetAllVaccineStatuses() []model.VaccineStatus {
	return []model.VaccineStatus{}
}
func (m *MockVaccineStatusSvc) GetVaccineStatusByID(id int) (model.VaccineStatus, error) {
	return model.VaccineStatus{}, nil
}
func (m *MockVaccineStatusSvc) DeleteVaccineStatusByID(id int) error {
	return nil
}
