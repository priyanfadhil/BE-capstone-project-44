package usecase

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type RepoMockSession struct {
	Mock mock.Mock
}

func (m *RepoMockSession) CreateSession(booking model.Session) error {
	arg := m.Mock.Called(booking)

	return arg.Error(0)
}
func (m *RepoMockSession) UpdateOneSessionByID(id int, booking model.Session) error {
	return nil
}
func (m *RepoMockSession) GetAllSessions() []model.Session {
	return []model.Session{}
}
func (m *RepoMockSession) GetOneSessionByID(id int) (booking model.Session, err error) {
	return model.Session{}, nil
}
func (m *RepoMockSession) DeleteSessionByID(id int) error {
	return nil
}
