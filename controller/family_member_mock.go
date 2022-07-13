package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockFamilyMemberSvc struct {
	Mock mock.Mock
}

func (m *MockFamilyMemberSvc) CreateFamilyMember(familymember model.FamilyMember) error {
	arg := m.Mock.Called(familymember)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockFamilyMemberSvc) UpdateFamilyMember(id int, familymember model.FamilyMember) error {
	return nil
}
func (m *MockFamilyMemberSvc) GetAllFamilyMembers() []model.FamilyMember {
	return []model.FamilyMember{}
}
func (m *MockFamilyMemberSvc) GetFamilyMemberByID(id int) (model.FamilyMember, error) {
	return model.FamilyMember{}, nil
}
func (m *MockFamilyMemberSvc) DeleteFamilyMemberByID(id int) error {
	return nil
}
