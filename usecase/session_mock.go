package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockSession struct {
	update func(id int, session model.Session) error
	create func(session model.Session) error
	delete func(id int) error
	getone func(id int) (session model.Session, err error)
}

func (r *RepoMockSession) CreateSession(session model.Session) error {
	return r.create(session)
}
func (r *RepoMockSession) GetAllSessions() []model.Session {
	panic("implemente")
}
func (r *RepoMockSession) GetOneSessionByID(id int) (session model.Session, err error) {
	return r.getone(id)
}
func (r *RepoMockSession) UpdateOneSessionByID(id int, session model.Session) error {
	return r.update(id, session)
}
func (r *RepoMockSession) DeleteSessionByID(id int) error {
	return r.delete(id)
}

func (r *RepoMockSession) GetOneSessionByEmail(email string) (session model.Session, err error) {
	panic("implemente")
}
