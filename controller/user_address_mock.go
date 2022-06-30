package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockUserAddressSvc struct {
	Mock mock.Mock
}

func (m *MockUserAddressSvc) CreateUserAddress(useraddress model.UserAddress) error {
	arg := m.Mock.Called(useraddress)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockUserAddressSvc) UpdateUserAddress(id int, useraddress model.UserAddress) error {
	return nil
}
func (m *MockUserAddressSvc) GetAllUserAddresses() []model.UserAddress {
	return []model.UserAddress{}
}
func (m *MockUserAddressSvc) GetUserAddressByID(id int) (model.UserAddress, error) {
	return model.UserAddress{}, nil
}
func (m *MockUserAddressSvc) DeleteUserAddressByID(id int) error {
	return nil
}
