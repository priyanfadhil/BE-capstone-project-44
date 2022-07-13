package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockAdminSvc struct {
	Mock mock.Mock
}

func (m *MockAdminSvc) CreateAdmin(admin model.Admin) error {
	arg := m.Mock.Called(admin)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockAdminSvc) UpdateAdmin(id int, admin model.Admin) error {
	return nil
}
func (m *MockAdminSvc) GetAllAdmins() []model.Admin {
	return []model.Admin{}
}
func (m *MockAdminSvc) GetAdminByID(id int) (model.Admin, error) {
	return model.Admin{}, nil
}
func (m *MockAdminSvc) DeleteAdminByID(id int) error {
	return nil
}
