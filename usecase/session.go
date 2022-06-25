package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcSession struct {
	c    config.Config
	repo domain.AdapterSessionRepository
}

func (s *svcSession) CreateSession(session model.Session) error {
	return s.repo.CreateSession(session)
}

func (s *svcSession) UpdateSession(id int, session model.Session) error {
	fmt.Println(id)
	return s.repo.UpdateOneSessionByID(id, session)
}

func (s *svcSession) GetAllSessions() []model.Session {
	return s.repo.GetAllSessions()
}

func (s *svcSession) GetSessionByID(id int) (model.Session, error) {
	return s.repo.GetOneSessionByID(id)
}

func (s *svcSession) DeleteSessionByID(id int) error {
	return s.repo.DeleteSessionByID(id)
}

func NewSession(repo domain.AdapterSessionRepository, c config.Config) domain.AdapterSession {
	return &svcSession{
		repo: repo,
		c:    c,
	}
}
