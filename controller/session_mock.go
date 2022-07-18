package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockSessionSvc struct {
	Mock mock.Mock
}

func (m *MockSessionSvc) CreateSession(session model.Session) error {
	arg := m.Mock.Called(session)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockSessionSvc) UpdateSession(id int, session model.Session) error {
	return nil
}
func (m *MockSessionSvc) GetAllSessions() []model.Session {
	return []model.Session{}
}
func (m *MockSessionSvc) GetSessionByID(id int) (model.Session, error) {
	return model.Session{}, nil
}
func (m *MockSessionSvc) DeleteSessionByID(id int) error {
	return nil
}
