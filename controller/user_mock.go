package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockUserSvc struct {
	Mock mock.Mock
}

func (m *MockUserSvc) CreateUser(email string, user model.User) error {
	arg := m.Mock.Called(user)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockUserSvc) UpdateUser(id, idToken int, user model.User) error {
	return nil
}
func (m *MockUserSvc) GetAllUsers() []model.User {
	return []model.User{}
}
func (m *MockUserSvc) GetUserByID(id int) (model.User, error) {
	return model.User{}, nil
}
func (m *MockUserSvc) DeleteUserByID(id int) error {
	return nil
}
func (m *MockUserSvc) LoginUser(email, password string) (string, int){
	return "", 0
}
