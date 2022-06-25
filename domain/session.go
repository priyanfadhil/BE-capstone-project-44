package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterSessionRepository interface {
	CreateSession(Session model.Session) error
	GetAllSessions() []model.Session
	GetOneSessionByID(id int) (Session model.Session, err error)
	UpdateOneSessionByID(id int, Session model.Session) error
	DeleteSessionByID(id int) error
}

// Use Case
type AdapterSession interface {
	CreateSession(Session model.Session) error
	UpdateSession(id int, Session model.Session) error
	GetAllSessions() []model.Session
	GetSessionByID(id int) (model.Session, error)
	DeleteSessionByID(id int) error
}
